# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 578J

Tested At: 31 Oct 22 16:40 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 268 day(s)\
Subject: CN=SHAKEN 578J, OU=SHAKEN, O=Call Tools Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/578J/8e374e10-eba0-45b0-b09a-b7086947dd7c.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8DCCApagAwIBAgIQeVbd7aKxW%2BT6O514gxNtWDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MjUxODM2NTNaFw0yMzA3MjUxODM2NTJaME0xCzAJBgNVBAYTAlVTMRcwFQYDVQQKEw5DYWxsIFRvb2xzIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNTc4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOyJ%2FwxL7DIqWDJK1%2FttiRpqWXBKSidvO4aGIGW64iTGJ%2Be2A5eXQw9EbgMXeji9Nur6sfjh4XMV7gSEFayDULujggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQU6xaKWcWmwT7CSI5kGM7vGtAHgowwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENTc4SjAKBggqhkjOPQQDAgNIADBFAiEA1LhZXDp8XJ0z39K%2B%2FDEsk8L26d3RG17jNSOHeVzUlPMCIFmncW%2Bxtt3Ibb58J6WP10JXQEbJDfkzfiIkZCa73cRO)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 31/10/2022 at 16:43:22