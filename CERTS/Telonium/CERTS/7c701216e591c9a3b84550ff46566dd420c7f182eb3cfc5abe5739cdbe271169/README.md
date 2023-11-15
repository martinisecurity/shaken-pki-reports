# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Telonium STI-CA Intermediate CA

Tested At: 15 Nov 23 18:10 UTC\
Initial Validity Period: 3650 day(s)\
Remaining Validity Period: 3399 day(s)\
Subject: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Issuer: CN=Telonium STI-CA Root1, O=Telonium, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIB6TCCAY6gAwIBAgIQN9QyShF6ZHqe6LZKCsjJMzAKBggqhkjOPQQDAjBAMQswCQYDVQQGEwJVUzERMA8GA1UECgwIVGVsb25pdW0xHjAcBgNVBAMMFVRlbG9uaXVtIFNUSS1DQSBSb290MTAeFw0yMzAzMDkxNTE4MTBaFw0zMzAzMDYxNTE4MTBaMEQxGDAWBgNVBAoTD1RlbG9uaXVtIFNUSS1DQTEoMCYGA1UEAxMfVGVsb25pdW0gU1RJLUNBIEludGVybWVkaWF0ZSBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMpDEHs3iKEkFRW30GFHzF5lyBSMPU8SDeFgZX2rSQ9u%2BAphn3bPCRsQTQcU0TSCnra5cC7vOEbSK%2Bck8ZAYE5mjZjBkMA4GA1UdDwEB%2FwQEAwIBBjASBgNVHRMBAf8ECDAGAQH%2FAgEAMB0GA1UdDgQWBBSqJLv%2FFHVAeS2Hb%2BgNQXfKu82IsDAfBgNVHSMEGDAWgBQQogXr4agct5Yi2Teu9uwT2hfOVzAKBggqhkjOPQQDAgNJADBGAiEAp8MR3633wQ2ltoGP0dB828TdGc3j1RyH7newO4k7VNwCIQD76dpGbQZHyCKqTXLdzYxNPlCry3IPcpHwZ3vviHccnA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ca_certificate_policies](../../ISSUES/e_atis_ca_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| [e_atis_ca_subject](../../ISSUES/e_atis_ca_subject/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_ca_crl_distribution](../../ISSUES/e_atis_ca_crl_distribution/README.md) | error | ATIS1000080 | STI Intermediate certificates shall contain a CRL Distribution Points extension |
| [e_atis_ca_subject_cn](../../ISSUES/e_atis_ca_subject_cn/README.md) | error | ATIS1000080 | The Common Name attribute shall include the text string "SHAKEN" |


Generated: 15 Nov 23 18:10 UTC