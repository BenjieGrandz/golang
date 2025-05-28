flowchart LR
    subgraph "Private Key Creation"
        A["RSA Key Generation
        • 2048 bits strength
        • Used for signing"]
    end
    
    subgraph "Certificate Template"
        B["Template Setup
        • Serial Number
        • Subject Info
        • Validity Period
        • Usage Flags"]
    end
    
    subgraph "Self-Signing Process"
        C["Certificate Creation
        • Sign template with private key
        • Create DER encoded cert"]
        
        D["PEM Encoding
        • Convert cert to PEM
        • Convert key to PEM"]
        
        E["TLS Certificate
        • Load PEM files
        • Return tls.Certificate"]
    end
    
    A --> B
    B --> C
    C --> D
    D --> E
    
    style A fill:#f9f,stroke:#333,color:#000
    style B fill:#bbf,stroke:#333,color:#000
    style C fill:#bfb,stroke:#333,color:#000
    style D fill:#fbf,stroke:#333,color:#000
    style E fill:#ffb,stroke:#333,color:#000