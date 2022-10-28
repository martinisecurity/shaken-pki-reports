# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 2c1fc9cdde6833c8f868003987fa38a40b1fc733
Tested At: 2022-10-28 19:21:48 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 6 day(s)\
Subject: CN=SHAKEN 625J, OU=SHAKEN, O=Victory Telecom Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/b7343686-5ed8-402c-89a3-8bf1a3d48975/1b6c455dcc3cd7586f99609b4a7291af.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9jCCApygAwIBAgIQamaEhhL49kHAMJ7kudCLCDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMDIzNTJaFw0yMjExMDIyMDIzNTFaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRWaWN0b3J5IFRlbGVjb20gSW5jLjEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNjI1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMnDPvAXs57ZPLTqSu2fp7yO95QjrIitRO4WFWR7g3r1CxDUrjqrh%2FtIjogdtT4u4lvcDnp%2FZqT7J6fqaY3vFnijggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUSpTShpUvU4p0uSEJbRjFqfo%2BrSkwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENjI1SjAKBggqhkjOPQQDAgNIADBFAiB2o%2F5lNDsmK4xGni5sMogw5sQUhpFpjIg19Af%2F5i6IyAIhALDMC3YMNY8AtGrXQK50zLjk4I8pbvaVdNP%2Bhu9YUIUh)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 19:22:10