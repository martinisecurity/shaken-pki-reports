# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 551G

Tested At: 31 Oct 22 20:24 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 3 day(s)\
Subject: CN=SHAKEN 551G, OU=SHAKEN, O=Brightlink Communications LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/551G/cda98320-b55d-4677-b2f3-771ab42f88f2.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FzCCAqWgAwIBAgIQR0XMaFBmaID%2BcOLDIvVcOzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMjAwMjVaFw0yMjExMDIyMjAwMjRaMFwxCzAJBgNVBAYTAlVTMSYwJAYDVQQKEx1CcmlnaHRsaW5rIENvbW11bmljYXRpb25zIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNTUxRzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDNE3SbfCokoh4FphJzkraz8BqGchAUBPFLbHho7VjHagSuJkPwQKVwo1vsXekf5pQfW4DoxJpLoeHn6F96FjgOjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUfvbumIn09UVXAYXs6MERqwzsQMEwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENTUxRzAKBggqhkjOPQQDAgNIADBFAiAnfzALYL8PbCuC8qEmaGR2Hq9y5YMJPrKTiL1qxJIJhwIhAOZEydIBpBJfRoCe511WZrfJMioCtFwXiambNAW0j8w7)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31/10/2022 at 20:32:42