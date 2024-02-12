# STIR/SHAKEN CA Ecosystem Compliance

## Certificate HD CARRIER LLC

Tested At: 12 Feb 24 16:53 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -195 day(s)\
Subject: L=Long Beach, ST=California, O=HD CARRIER LLC, C=US, CN=HD CARRIER LLC\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://d64db847f381fcb974ab1b6150b49a91.cxstatic.com/f/e7025f18-eb5d-4cac-80ca-c69254cbbb3c

[View certificate details](https://understandingwebpki.com/?cert=MIICpDCCAiqgAwIBAgIIagt%2BPC1Y8lIwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIzMDcwMjAwMDAwMFoXDTIzMDczMTIzNTk1OVowaTEXMBUGA1UEAwwOSEQgQ0FSUklFUiBMTEMxCzAJBgNVBAYTAlVTMRcwFQYDVQQKDA5IRCBDQVJSSUVSIExMQzETMBEGA1UECAwKQ2FsaWZvcm5pYTETMBEGA1UEBwwKTG9uZyBCZWFjaDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD6OiJaXqKtjOJcUCGUsJi90kzz2iY2aCBIAW1X%2FoMtbYsiqSdk%2FLuMreSL%2BPXH0xKuVDvXHrrLa3XoNrnGuhC2jgZUwgZIwFgYIKwYBBQUHARoECjAIoAYWBDMyMUowDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFE7i5D9%2BbvYDihP%2BvrbbpFIS%2Fw6FMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNoADBlAjEA0QEktZ3c3cD%2BJhu6DJD1hHv2146a0qUz2386v19qVjxoLiberF9RJIFWYq8BuCo2AjBDyf4A1eVZGkcASLECOA%2B%2FvSy0qaUrB7Y1xd7%2FxlTYAviMgmRWisARXflM8UN4%2Frw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 63 |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'HD CARRIER LLC' does not contain 'SHAKEN' |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 321J', but common name is 'HD CARRIER LLC' |


Generated: 12 Feb 24 17:02 UTC