package viewcontext

import (
	"fmt"

	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/kv/optimism"
	"github.com/iotaledger/wasp/packages/kv/subrealm"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/vm/core/blob"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/iotaledger/wasp/packages/vm/processors"
	"golang.org/x/xerrors"
)

type viewcontext struct {
	processors  *processors.ProcessorCache
	stateReader state.OptimisticStateReader
	chainID     coretypes.ChainID
	log         *logger.Logger
}

func NewFromChain(chain chain.ChainCore) *viewcontext {
	return New(*chain.ID(), chain.GetStateReader(), chain.Processors(), chain.Log().Named("view"))
}

func New(chainID coretypes.ChainID, stateReader state.OptimisticStateReader, proc *processors.ProcessorCache, log *logger.Logger) *viewcontext {
	return &viewcontext{
		processors:  proc,
		stateReader: stateReader,
		chainID:     chainID,
		log:         log,
	}
}

// CallView in viewcontext implements own panic catcher.
func (v *viewcontext) CallView(contractHname coretypes.Hname, epCode coretypes.Hname, params dict.Dict) (dict.Dict, error) {
	var ret dict.Dict
	var err error
	func() {
		defer func() {
			r := recover()
			if r == nil {
				return
			}
			ret = nil
			switch err1 := r.(type) {
			case *kv.DBError:
				v.log.Panicf("DB error: %v", err1)
			case *optimism.ErrorStateInvalidated:
				err = err1
			default:
				err = xerrors.Errorf("viewcontext: panic in VM: %w", err1)
			}
		}()
		ret, err = v.callView(contractHname, epCode, params)
	}()
	return ret, err
}

func (v *viewcontext) callView(contractHname coretypes.Hname, epCode coretypes.Hname, params dict.Dict) (dict.Dict, error) {
	var err error
	contractRecord, err := root.FindContract(contractStateSubpartition(v.stateReader.KVStoreReader(), root.Interface.Hname()), contractHname)
	if err != nil {
		return nil, fmt.Errorf("failed to find contract %s: %v", contractHname, err)
	}
	proc, err := v.processors.GetOrCreateProcessor(contractRecord, func(programHash hashing.HashValue) (string, []byte, error) {
		if vmtype, ok := processors.GetBuiltinProcessorType(programHash); ok {
			return vmtype, nil, nil
		}
		return blob.LocateProgram(contractStateSubpartition(v.stateReader.KVStoreReader(), blob.Interface.Hname()), programHash)
	})
	if err != nil {
		return nil, err
	}

	ep, ok := proc.GetEntryPoint(epCode)
	if !ok {
		return nil, fmt.Errorf("%s: can't find entry point '%s'", proc.GetDescription(), epCode.String())
	}

	if !ep.IsView() {
		return nil, fmt.Errorf("only view entry point can be called in this context")
	}
	return ep.Call(newSandboxView(v, contractHname, params))
}

func contractStateSubpartition(state kv.KVStoreReader, contractHname coretypes.Hname) kv.KVStoreReader {
	return subrealm.NewReadOnly(state, kv.Key(contractHname.Bytes()))
}

func (v *viewcontext) Infof(format string, params ...interface{}) {
	v.log.Infof(format, params...)
}

func (v *viewcontext) Debugf(format string, params ...interface{}) {
	v.log.Debugf(format, params...)
}

func (v *viewcontext) Panicf(format string, params ...interface{}) {
	v.log.Panicf(format, params...)
}
