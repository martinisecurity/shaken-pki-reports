# STIR/SHAKEN CA Ecosystem Compliance

## Certificate HD CARRIER LLC

Tested At: 04 Oct 24 16:16 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -446 day(s)\
Subject: L=Long Beach, ST=California, O=HD CARRIER LLC, C=US, CN=HD CARRIER LLC\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://d64db847f381fcb974ab1b6150b49a91.cxstatic.com/f/20f8bca4-d82b-49d7-b419-ff8a42d2dfbb

[View certificate details](https://x509.io/?cert=MIICpTCCAiqgAwIBAgIIFbHOFgS9nzMwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIzMDYxNjAwMDAwMFoXDTIzMDcxNTIzNTk1OVowaTEXMBUGA1UEAwwOSEQgQ0FSUklFUiBMTEMxCzAJBgNVBAYTAlVTMRcwFQYDVQQKDA5IRCBDQVJSSUVSIExMQzETMBEGA1UECAwKQ2FsaWZvcm5pYTETMBEGA1UEBwwKTG9uZyBCZWFjaDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD6OiJaXqKtjOJcUCGUsJi90kzz2iY2aCBIAW1X%2FoMtbYsiqSdk%2FLuMreSL%2BPXH0xKuVDvXHrrLa3XoNrnGuhC2jgZUwgZIwFgYIKwYBBQUHARoECjAIoAYWBDMyMUowDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFE7i5D9%2BbvYDihP%2BvrbbpFIS%2Fw6FMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNpADBmAjEAs6Tf8YeeGTqtfhLQ4pJhubmA4zrijmfnABuieAnm8e7lZa%2FZGDcT2dVvjB9B63cbAjEAvi0BC8u29eT69GptmA8rI%2Bva0CdD9oPdA%2FgDToOvjomgvA0s0pB6nkm1N4qGrY09)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 321J', but common name is 'HD CARRIER LLC' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 61 |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'HD CARRIER LLC' does not contain 'SHAKEN' |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |


Generated: 04 Oct 24 16:29 UTC