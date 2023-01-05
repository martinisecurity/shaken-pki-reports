# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 093K

Tested At: 05 Jan 23 18:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -57 day(s)\
Subject: CN=SHAKEN 093K, OU=SHAKEN, O=Skywave Wireless Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/093K/574481eb-e555-46a5-82e0-cb1e97ac0603.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9jCCApygAwIBAgIQSigT%2BiiFxNMhir7mYe7QZDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTAxODMzMjFaFw0yMjExMDkxODMzMjBaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRTa3l3YXZlIFdpcmVsZXNzIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDkzSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABP%2FtxjBHRGf%2B9Yy7YQ6JUHEne1qkWuvO4c6ydrohV48zfCy3W1aDxJ7pRIbXQdAhsPPASnnDm15jpEZjjeJhQB%2BjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUfUonLK28qo1shQQYFkFR3ZTMskAwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDkzSzAKBggqhkjOPQQDAgNIADBFAiEA3Asf5ADbkm57m%2BBY7eIfwA3zASpChZRlSaaFxNtxcpICICZIvVGl4%2FQDtLUTCyDaXZghIlfQbcnoChDi5lqU7sgQ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 05 Jan 23 18:35 UTC