# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate d33fcbf0b9ea652abf3b0deedaf23a6099b9fa64
Tested At: 2022-10-28 19:22:05 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: CN=SHAKEN 089K, OU=SHAKEN, O=Logista Solutions, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/5f7c135d-caf3-4661-abb1-d1720e7872e2/06bce0219cbd2420e11523954915f6ef.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9DCCApmgAwIBAgIQb0k9qNvCEinp0RI7WU4YGTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYxODM2MTJaFw0yMjExMDIxODM2MTFaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFMb2dpc3RhIFNvbHV0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDg5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNVIqJWJA7eqFp%2Bqb1lqcbz%2FfZNXvbaoOPVLjfUWd5sGkpEjz7zEcpS%2F%2FD5q6ukEMZ3QqvM7kBYgbODn5U2LLRmjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUxDqy6DQqb85Q7KJnqDCoGLyx2lIwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDg5SzAKBggqhkjOPQQDAgNJADBGAiEAhPuxjXncivp7nitF2XcZ%2BVCnlyvbthaylODOMZ4x9IICIQCzLMVtDcmYXdU473ezwbsL7chMUtBnZHCjvutRfgwvnA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 28/10/2022 at 19:22:10