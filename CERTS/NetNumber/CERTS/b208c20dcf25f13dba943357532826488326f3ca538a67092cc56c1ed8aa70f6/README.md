# STIR/SHAKEN CA Ecosystem Compliance

## Certificate HD CARRIER LLC

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -14 day(s)\
Subject: L=Long Beach, ST=California, O=HD CARRIER LLC, C=US, CN=HD CARRIER LLC\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://d64db847f381fcb974ab1b6150b49a91.cxstatic.com/f/adbd7223-7185-4f25-8ba0-43ae9208b95f

[View certificate details](https://understandingwebpki.com/?cert=MIICozCCAiqgAwIBAgIINgn94l%2F%2FgqQwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIzMTAwMjAwMDAwMFoXDTIzMTAzMTIzNTk1OVowaTEXMBUGA1UEAwwOSEQgQ0FSUklFUiBMTEMxCzAJBgNVBAYTAlVTMRcwFQYDVQQKDA5IRCBDQVJSSUVSIExMQzETMBEGA1UECAwKQ2FsaWZvcm5pYTETMBEGA1UEBwwKTG9uZyBCZWFjaDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD6OiJaXqKtjOJcUCGUsJi90kzz2iY2aCBIAW1X%2FoMtbYsiqSdk%2FLuMreSL%2BPXH0xKuVDvXHrrLa3XoNrnGuhC2jgZUwgZIwFgYIKwYBBQUHARoECjAIoAYWBDMyMUowDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFE7i5D9%2BbvYDihP%2BvrbbpFIS%2Fw6FMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNnADBkAjBuQKLqt3AY93pihvF%2FLi9s02W874NvqfE2AwIo0XbMelGIdXfD8tJzZonnfZgwEPsCMBCarwuMBPCcJ9QuBcPwIdqsrLVg6ymi8dJYxPzxy68VXiB2f2BVBk7laxSAMVxTBw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 321J' |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |


Generated: 15 Nov 23 18:10 UTC