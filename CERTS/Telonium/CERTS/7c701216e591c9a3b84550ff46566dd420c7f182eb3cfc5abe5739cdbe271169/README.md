# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Telonium STI-CA Intermediate CA

Tested At: 28 Nov 23 20:19 UTC\
Initial Validity Period: 3650 day(s)\
Remaining Validity Period: 3386 day(s)\
Subject: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Issuer: CN=Telonium STI-CA Root1, O=Telonium, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIB6TCCAY6gAwIBAgIQN9QyShF6ZHqe6LZKCsjJMzAKBggqhkjOPQQDAjBAMQswCQYDVQQGEwJVUzERMA8GA1UECgwIVGVsb25pdW0xHjAcBgNVBAMMFVRlbG9uaXVtIFNUSS1DQSBSb290MTAeFw0yMzAzMDkxNTE4MTBaFw0zMzAzMDYxNTE4MTBaMEQxGDAWBgNVBAoTD1RlbG9uaXVtIFNUSS1DQTEoMCYGA1UEAxMfVGVsb25pdW0gU1RJLUNBIEludGVybWVkaWF0ZSBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMpDEHs3iKEkFRW30GFHzF5lyBSMPU8SDeFgZX2rSQ9u%2BAphn3bPCRsQTQcU0TSCnra5cC7vOEbSK%2Bck8ZAYE5mjZjBkMA4GA1UdDwEB%2FwQEAwIBBjASBgNVHRMBAf8ECDAGAQH%2FAgEAMB0GA1UdDgQWBBSqJLv%2FFHVAeS2Hb%2BgNQXfKu82IsDAfBgNVHSMEGDAWgBQQogXr4agct5Yi2Teu9uwT2hfOVzAKBggqhkjOPQQDAgNJADBGAiEAp8MR3633wQ2ltoGP0dB828TdGc3j1RyH7newO4k7VNwCIQD76dpGbQZHyCKqTXLdzYxNPlCry3IPcpHwZ3vviHccnA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_ca](../../ISSUES/e_atis_subject_cn_ca/README.md) | error | ATIS1000080 | Common Name attribute 'Telonium STI-CA Intermediate CA' does not contain 'SHAKEN' |
| [e_atis_subject_c_iso_ca](../../ISSUES/e_atis_subject_c_iso_ca/README.md) | error | ATIS1000080 | Subject MUST be present and MUST contain exactly one value for Country (C=). |
| [e_atis_ext_certificate_policies_ca](../../ISSUES/e_atis_ext_certificate_policies_ca/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [e_shaken_certificate_policies_id_ca](../../ISSUES/e_shaken_certificate_policies_id_ca/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_atis_ext_crl_distribution_ca](../../ISSUES/e_atis_ext_crl_distribution_ca/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_subject_dn_ca](../../ISSUES/e_atis_subject_dn_ca/README.md) | error | ATIS1000080 | Subject DN does not contain a Country (C=) attribute |


Generated: 28 Nov 23 20:21 UTC