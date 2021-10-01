/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

module.exports = {
    // By default, Docusaurus generates a sidebar from the docs folder structure
    //tutorialSidebar: [{type: 'autogenerated', dirName: '.'}],

    // But you can create a sidebar manually
    tutorialSidebar: [
        {
            type: 'doc',
            label: 'Overview',
            id: 'overview',
        },
        {
            type: 'category',
            label: 'Core Concepts',
            items: [
                {
                    type: 'doc',
                    label: 'Smart Contracts',
                    id: 'guide/core_concepts/smart-contracts',
                },
                {
                    type: 'doc',
                    label: 'ISCP',
                    id: 'guide/core_concepts/iscp',
                },
                {
                    type: 'doc',
                    label: 'ISCP Architecture',
                    id: 'guide/core_concepts/iscp-architecture',
                },
                {
                    type: 'doc',
                    label: 'Committees and Validators',
                    id: 'guide/core_concepts/validators',
                },
                {
                    type: 'doc',
                    label: 'Consensus',
                    id: 'guide/core_concepts/consensus',
                },
                {
                    type: 'doc',
                    label: 'State, Transitions and State Anchoring',
                    id: 'guide/core_concepts/states',
                },
                {
                    type: 'doc',
                    label: 'Anatomy of a Smart Contract',
                    id: 'guide/core_concepts/smart-contract-anatomy'
                },
                {
                    type: 'doc',
                    label: 'Sandbox Interface',
                    id: 'guide/core_concepts/sandbox'
                },
                {
                    type: 'category',
                    label: 'Core Contracts',
                    items: [
                        {
                            type: 'doc',
                            label: 'Overview',
                            id: 'guide/core_concepts/core_contracts/overview',
                        },
                        {
                            type: 'doc',
                            label: 'Root',
                            id: 'guide/core_concepts/core_contracts/root',
                        },
                        {
                            type: 'doc',
                            label: 'Default',
                            id: 'guide/core_concepts/core_contracts/default',
                        },
                        {
                            type: 'doc',
                            label: 'Accounts',
                            id: 'guide/core_concepts/core_contracts/accounts',
                        },
                        {
                            type: 'doc',
                            label: 'Blob',
                            id: 'guide/core_concepts/core_contracts/blob',
                        },
                        {
                            type: 'doc',
                            label: 'Blocklog',
                            id: 'guide/core_concepts/core_contracts/blocklog',
                        },
                    ],
                },
                {
                    type: 'category',
                    label: 'Accounts',
                    items: [
                        {
                            type: 'doc',
                            label: 'How Accounts Work',
                            id: 'guide/core_concepts/accounts/how-accounts-work',
                        },
                        {
                            type: 'doc',
                            label: 'How To Deposit To a Chain',
                            id: 'guide/core_concepts/accounts/how-to-deposit-to-a-chain',
                        },
                        {
                            type: 'doc',
                            label: 'How To Withdraw From a Chain',
                            id: 'guide/core_concepts/accounts/how-to-withdraw-from-a-chain',
                        },
                        {
                            type: 'doc',
                            label: 'View Account Balances',
                            id: 'guide/core_concepts/accounts/view-account-balances',
                        },
                        {
                            type: 'doc',
                            label: 'The Common Account',
                            id: 'guide/core_concepts/accounts/the-common-account',
                        },
                    ]
                },
                {
                    type: 'category',
                    label: 'Interacting With Smart Contracts',
                    items: [
                        {
                            type: 'doc',
                            label: 'On-ledger Requests',
                            id: 'guide/core_concepts/smartcontract-interaction/on-ledger-requests',
                        },
                        {
                            type: 'doc',
                            label: 'Off-ledger Requests',
                            id: 'guide/core_concepts/smartcontract-interaction/off-ledger-requests',
                        }
                    ]
                },
                {
                    type: 'category',
                    label: 'Types of VMs/Languages',
                    items: [
                        {
                            type: 'doc',
                            label: 'How ISCP Works As a Language/VM Agnostic Platform',
                            id: 'guide/core_concepts/vm-types/iscp-vm',
                        },
                        {
                            type: 'category',
                            label: 'Rust/Wasm Based Smart Contracts',
                            items: [
                                {
                                    type: 'doc',
                                    label: 'Why and What Does It Look Like?',
                                    id: 'guide/core_concepts/vm-types/rust-wasm',
                                },
                                {
                                    type: 'doc',
                                    label: 'Solidity/EVM Based Smart Contracts',
                                    id: 'guide/core_concepts/vm-types/rust-wasm',
                                },
                            ]
                        },
                        {
                            type: 'category',
                            label: 'Solidity/EVM Based Smart Contracts',
                            items: [
                                {
                                    type: 'doc',
                                    label: 'Why and What Does It Look Like?',
                                    id: 'guide/core_concepts/vm-types/evm',
                                },
                                {
                                    type: 'doc',
                                    label: 'Solidity/EVM Based Smart Contracts',
                                    id: 'guide/core_concepts/vm-types/evm',
                                },
                            ]
                        }
                    ],
                },
                {
                    type: 'category',
                    label: 'Solo',
                    items: [
                        {
                            type: 'doc',
                            label: 'Solo',
                            id: 'guide/solo/what-is-solo',
                        },
                        {
                            type: 'doc',
                            label: 'First Example',
                            id: 'guide/solo/first-example',
                        },
                        {
                            type: 'doc',
                            label: 'Tokens and the UTXO Ledger',
                            id: 'guide/solo/tokens-and-utxos',
                        },
                        {
                            type: 'doc',
                            label: 'Invoking a Smart Contract',
                            id: 'guide/solo/invoking-sc',
                        },
                        {
                            type: 'doc',
                            label: 'Calling a View',
                            id: 'guide/solo/view-sc',
                        },
                        {
                            type: 'doc',
                            label: 'Error Handling',
                            id: 'guide/solo/error-handling',
                        },
                        {
                            type: 'doc',
                            label: 'Account Balances',
                            id: 'guide/solo/balances'
                        },
                        {
                            type: 'doc',
                            label: 'Sending tokens to a Smart Contract',
                            id: 'guide/solo/sending-funds'
                        },
                        {
                            type: 'doc',
                            label: 'Return of tokens in case of failure',
                            id: 'guide/solo/reimbursed-funds'
                        },
                        {
                            type: 'doc',
                            label: 'Sending tokens from ISCP to the Tangle',
                            id: 'guide/solo/sending-funds-sc'
                        },
                    ]
                }
            ],
        },
        {
            type: 'category',
            label: 'ISCP Chains and Nodes',
            items: [
                {
                    type: 'doc',
                    label: 'Running a Node',
                    id: 'guide/chains_and_nodes/running-a-node',
                },
                {
                    type: 'doc',
                    label: 'Configuring wasp-cli',
                    id: 'guide/chains_and_nodes/wasp-cli',
                },
                {
                    type: 'doc',
                    label: 'Setting Up a Chain',
                    id: 'guide/chains_and_nodes/setting-up-a-chain',
                },
                {
                    type: 'doc',
                    label: 'Chain Management',
                    id: 'guide/chains_and_nodes/chain-management',
                },
            ]
        },
        {
            type: 'category',
            label: 'Rust/WASM',
            items: [
                {
                    type: 'doc',
                    label: 'Introduction',
                    id: 'guide/rust_wasm/introduction',
                },
            ]
        },
        {
            type: 'category',
            label: 'EVM',
            items: [
                {
                    type: 'doc',
                    label: 'Introduction',
                    id: 'guide/evm/introduction',
                },
                {
                    type: 'doc',
                    label: 'Limitations',
                    id: 'guide/evm/limitations',
                },
                {
                    type: 'doc',
                    label: 'Creating a chain',
                    id: 'guide/evm/create-chain',
                },
                {
                    type: 'doc',
                    label: 'Example contract',
                    id: 'guide/evm/example',
                },
                {
                    type: 'doc',
                    label: 'Tooling',
                    id: 'guide/evm/tooling',
                }
            ]
        },
        {
            type: 'category',
            label: 'Example projects',
            items: [
                {
                    type: 'doc',
                    label: 'Fair Roulette',
                    id: 'guide/example_projects/fair_roulette',
                },
            ]
        },
        {
            type: 'doc',
            label: 'Contribute',
            id: 'contribute',
        }
    ],
};
