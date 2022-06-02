module.exports = {
  theme: 'cosmos',
  title: 'Fortress Documentation',
  locales: {
    '/': {
      lang: 'en-US'
    },
  },
  markdown: {
    extendMarkdown: (md) => {
      md.use(require("markdown-it-katex"));
    },
  },
  head: [
    [
      "link",
      {
        rel: "stylesheet",
        href:
          "https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.5.1/katex.min.css",
      },
    ],
    [
      "link",
      {
        rel: "stylesheet",
        href:
          "https://cdn.jsdelivr.net/github-markdown-css/2.2.1/github-markdown.css",
      },
    ],
  ],
  base: process.env.VUEPRESS_BASE || '/',
  plugins: [
    'vuepress-plugin-element-tabs'
  ],
  head: [
    // ['link', { rel: "apple-touch-icon", sizes: "180x180", href: "/apple-touch-icon.png" }],
    ['link', { rel: "icon", type: "image/png", sizes: "32x32", href: "/favicon32.png" }],
    ['link', { rel: "icon", type: "image/png", sizes: "16x16", href: "/favicon16.png" }],
    ['link', { rel: "manifest", href: "/site.webmanifest" }],
    ['meta', { name: "msapplication-TileColor", content: "#2e3148" }],
    ['meta', { name: "theme-color", content: "#ffffff" }],
    ['link', { rel: "icon", type: "image/svg+xml", href: "/favicon.svg" }],
    // ['link', { rel: "apple-touch-icon-precomposed", href: "/apple-touch-icon-precomposed.png" }],
  ],
  themeConfig: {
    repo: 'Karan-3108/fortress',
    docsRepo: 'Karan-3108/fortress',
    docsBranch: 'main',
    docsDir: 'docs',
    editLinks: true,
    custom: true,
    project: {
      name: 'Fortress',
      denom: 'Fortress',
      ticker: 'FORTRESS',
      binary: 'fortressd',
      testnet_denom: 'tFortress',
      testnet_ticker: 'tFORTRESS',
      rpc_url: 'https://eth.bd.fortress.org:8545',
      rpc_url_testnet: 'https://eth.bd.fortress.dev:8545',
      rpc_url_local: 'http://localhost:8545/',
      chain_id: '9001',
      testnet_chain_id: '9000',
      latest_version: 'v3.0.0-beta1',
      version_number: '2',
      testnet_version_number: '4',
      testnet_evm_explorer_url: 'https://evm.fortress.dev',
      evm_explorer_url: 'https://evm.fortress.org',
      testnet_cosmos_explorer_url: 'https://explorer.fortress.dev/',
      cosmos_explorer_url: 'https://www.mintscan.io/fortress',
    },
    logo: {
      src: '/fortress-black.svg',
    },
    algolia: {
      id: 'BH4D9OD16A',
      key: 'a5d55fe5f540cc3bd28fa2c72f2b5bd8',
      index: 'fortress'
    },
    topbar: {
      banner: false
    },
    sidebar: {
      auto: false,
      nav: [
        {
          title: 'About Fortress',
          children: [
            {
              title: 'Introduction',
              directory: true,
              path: '/about/intro'
            },
            {
              title: 'Fortress Ecosystem',
              path: 'https://fortress.space/'
            },
            {
              title: 'Awesome Fortress',
              path: 'https://github.com/Karan-3108/awesome'
            },
          ]
        },
        {
          title: 'For Users',
          children: [
            {
              title: 'Basic Concepts',
              directory: true,
              path: '/users/basics'
            },
            {
              title: 'Digital Wallets',
              directory: true,
              path: '/users/wallets'
            },
            {
              title: 'Account Keys',
              directory: true,
              path: '/users/keys'
            },
            {
              title: 'Technical Concepts',
              directory: true,
              path: '/users/technical_concepts'
            },
          ]
        },
        {
          title: 'For dApp Devs',
          children: [
            {
              title: 'Overview',
              directory: false,
              path: '/developers/overview'
            },
            {
              title: 'Quick Connect',
              directory: false,
              path: '/developers/connect'
            },
            {
              title: 'Fortress Clients',
              directory: false,
              path: '/developers/clients'
            },
            {
              title: 'Testnet Faucet',
              directory: false,
              path: '/developers/faucet'
            },
            {
              title: 'Ethereum Tooling',
              directory: true,
              path: '/developers/tools'
            },
            {
              title: 'Localnet',
              directory: true,
              path: '/developers/localnet'
            },
            {
              title: 'Libraries',
              directory: true,
              path: '/developers/libraries'
            },
            {
              title: 'Ethereum JSON-RPC',
              directory: true,
              path: '/developers/json-rpc'
            },
            {
              title: 'Cosmos gRPC & REST',
              path: 'https://api.fortress.dev/'
            },
            {
              title: 'Tendermint RPC',
              path: 'https://docs.tendermint.com/v0.34/rpc/'
            },
          ]
        },
        {
          title: 'For Protocol Devs',
          children: [
            {
              title: 'Module Specifications',
              directory: true,
              path: '/modules'
            },
            {
              title: 'Fortress Go API',
              path: 'https://pkg.go.dev/github.com/Karan-3108/fortress'
            },
            {
              title: 'Ethermint Library Go API',
              path: 'https://pkg.go.dev/github.com/Karan-3108/ethermint'
            },
            {
              title: 'Fortress Protobuf',
              directory: false,
              path: '/protocol/proto-docs'
            }
          ]
        },
        {
          title: 'For Validators',
          children: [
            {
              title: 'Validators Overview',
              directory: false,
              path: '/validators/overview'
            },
            {
              title: 'Setup & Configuration',
              directory: true,
              path: '/validators/setup'
            },
            {
              title: 'Installation & Quick Start',
              directory: true,
              path: '/validators/quickstart'
            },
            {
              title: 'Join Testnet',
              directory: false,
              path: '/validators/testnet'
            },
            {
              title: 'Join Mainnet',
              directory: false,
              path: '/validators/mainnet'
            },
            {
              title: 'Security',
              directory: true,
              path: '/validators/security'
            },
            {
              title: 'Software Upgrade Guide',
              directory: true,
              path: '/validators/upgrades'
            },
            {
              title: 'Fortress Governance',
              directory: true,
              path: '/validators/governance'
            },
            {
              title: 'FAQ',
              directory: false,
              path: '/validators/faq'
            }
          ]
        },
        {
          title: 'Block Explorers',
          children: [
            {
              title: 'Block Explorers',
              path: '/developers/explorers'
            },
            {
              title: 'Blockscout (EVM)',
              path: 'https://evm.fortress.org'
            },
            {
              title: 'Mintscan (Cosmos)',
              path: 'https://www.mintscan.io/fortress/'
            },
          ]
        },
      ]
    },
    gutter: {
      title: 'Help & Support',
      chat: {
        title: 'Discord Channel',
        text: 'Chat with Fortress users and team on Discord.',
        url: 'https://discord.gg/fortress',
        bg: 'linear-gradient(103.75deg, #1B1E36 0%, #22253F 100%)'
      },
      forum: {
        title: 'Commonwealth Forum',
        text: 'Join the Fortress Commonwealth forum',
        url: 'https://commonwealth.im/fortress', 
        bg: 'linear-gradient(221.79deg, #3D6B99 -1.08%, #336699 95.88%)',
      },
      github: {
        title: 'Found an Issue?',
        text: 'Help us improve this page by suggesting edits on GitHub.',
        bg: '#F8F9FC'
      }
    },
    footer: {
      logo: '/fortress-black.svg',
      textLink: {
        text: "fortress.org",
        url: 'https://fortress.org'
      },
      services: [
        {
          service: "github",
          url: 'https://github.com/Karan-3108/fortress'
        },
        {
          service: "twitter",
          url: "https://twitter.com/FortressOrg",
        },
        {
          service: "telegram",
          url: "https://t.me/FortressOrg",
        },
        {
          service: "linkedin",
          url: "https://www.linkedin.com/company/Karan-3108-finance/",
        },
        {
          service: "medium",
          url: "https://fortress.blog/",
        },
      ],
      smallprint: 'This website is maintained by Karan-3108 Labs Ltd.',
      links: [{
        title: 'Ecosystem Documentation',
        children: [
        {
          title: 'Cosmos SDK Docs',
          url: 'https://docs.cosmos.network'
        },
        {
          title: 'Ethereum Docs',
          url: 'https://ethereum.org/developers'
        },
        {
          title: 'Tendermint Core Docs',
          url: 'https://docs.tendermint.com'
        }
        ]
      },
      {
        title: 'Community',
        children: [
          {
            title: 'Fortress Discord Community',
            url: 'https://discord.gg/fortress'
          },
          {
            title: 'Fortress Commonwealth Forum',
            url: 'https://commonwealth.im/fortress'
          },
        ]
      },
      {
        title: 'Fortress',
        children: [
          {
            title: 'Jobs at Fortress',
            url: 'https://Karan-3108.notion.site/'
          }
        ]
      }
      ]
    },
    versions: [
      {
        "label": "main",
        "key": "main"
      },
    ],
  }
};
