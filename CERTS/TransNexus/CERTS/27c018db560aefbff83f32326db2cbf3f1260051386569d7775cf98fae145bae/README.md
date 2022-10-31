# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 345J

Tested At: 31 Oct 22 18:16 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 21 day(s)\
Subject: CN=SHAKEN 345J, OU=SHAKEN, O=Ooma Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/345J/84bd095a-afff-4ef9-bb30-3061e3e407cb.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6TCCApCgAwIBAgIQQlOVLJtn9rLUHE6xwsuUMTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjIwNDA0NDZaFw0yMjExMjEwNDA0NDVaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhPb21hIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzQ1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABE4j01v0J4mKojaCST%2FJBS2JoBkj2j0ab3DlJPYYWMIimqwfID94IkgwA17EINByib1y%2Fex8bD7wkKLX6LakhvyjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUh5bMbFF2ksnzvO18J0NB8FoalAwwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzQ1SjAKBggqhkjOPQQDAgNHADBEAiBCM9WhNm2UOA6wyjYO5AEW1lWMG555Ev77mx217qHBZQIgWhXnMYfN9IKNJkKi%2F0yWbKdhZzlk9XnCxPx4y6675dw%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 31/10/2022 at 18:25:03