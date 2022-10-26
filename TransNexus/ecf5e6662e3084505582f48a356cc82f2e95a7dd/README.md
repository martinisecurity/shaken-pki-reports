# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate ecf5e6662e3084505582f48a356cc82f2e95a7dd
Tested At: 2022-10-26 23:14:36 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 26 day(s)\
Subject: CN=SHAKEN 345J, OU=SHAKEN, O=Ooma Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/345J/ddddd245-79a7-4779-a067-cc56739bf26b.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6jCCApCgAwIBAgIQTtltzykJ0rDqxWjLTFADkjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjIwNDQ0MDRaFw0yMjExMjEwNDQ0MDNaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhPb21hIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzQ1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD6Kf4C4LhJ2zSkUACdrNiALYVjCJY2K%2BAzotOc6BXZMaBHD65A78pI%2F1huGRMGETEtSq%2FX4%2B4AtLwXqaNRL3QGjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUQDDTXxnoiLSADuU48Id%2BChNA%2B88wHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzQ1SjAKBggqhkjOPQQDAgNIADBFAiAhCWs%2BKBo9k1QaYjHSsjQ9ApezI4umYPbA7qrSS1m8mAIhANnY4AHsWosTOX6YQgQmLgFE9SRPsxYM%2BjC2mafzDJBs)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |

Generated: 26/10/2022 at 23:14:41