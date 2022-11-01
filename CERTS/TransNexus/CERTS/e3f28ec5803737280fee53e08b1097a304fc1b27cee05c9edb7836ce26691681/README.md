# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 366G

Tested At: 01 Nov 22 20:29 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN 366G, OU=SHAKEN, O=USA Digital, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/e9497931-96b3-4ff8-8306-1b1273847a4d/664da3349fd259a9a51e89e03db614e0.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7TCCApOgAwIBAgIQd2EW8VaeHIfPNYsPd8a%2FJzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMDI1MDBaFw0yMjExMDIyMDI0NTlaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtVU0EgRGlnaXRhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzY2RzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKqkajTpLbdTDMPw6AspMklXxm5IJ8RM3WWvkt7Rcqpw6CH%2BikBj6ZLwJlpQ2TWPYsTK3iTi2Hx7aG1qZPwOkG2jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUXNgXXkNsoVFkqNT7YL08j7qUKyEwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzY2RzAKBggqhkjOPQQDAgNIADBFAiArdiezT0wOs1UA8AYd0HuBcQTuswjN2aw9sfdsEi1VwgIhAIirV%2FQKfScu%2FTpBtDm2V8Z8jWyaGvvOhuVigq2wUuok)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 01/11/2022 at 20:31:14