# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Definitely SHAKEN 709J but not SHAKEN 0007

Tested At: 02 Nov 22 15:33 UTC\
Initial Validity Period: 1 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=Definitely SHAKEN 709J but not SHAKEN 0007, OU=Trolling division, O=Low Latency Communications LLC - Self Reported, L=Birmingham Metro, ST=Alabama, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://sketchy.gay/shaken/llc-cert-3.pem

View: [View certificate details](https://understandingwebpki.com/?cert=MIIDfTCCAyOgAwIBAgIQRYxwkJA3sxYwrh%2Bey4bg%2FDAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjExMDExOTU5MzhaFw0yMjExMDIxODEzMzNaMIHEMQswCQYDVQQGEwJVUzEQMA4GA1UECAwHQWxhYmFtYTEZMBcGA1UEBwwQQmlybWluZ2hhbSBNZXRybzE3MDUGA1UECgwuTG93IExhdGVuY3kgQ29tbXVuaWNhdGlvbnMgTExDIC0gU2VsZiBSZXBvcnRlZDEaMBgGA1UECwwRVHJvbGxpbmcgZGl2aXNpb24xMzAxBgNVBAMMKkRlZmluaXRlbHkgU0hBS0VOIDcwOUogYnV0IG5vdCBTSEFLRU4gMDAwNzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABBQWvd1OijCPfV4Nz9DwErKzIBU9%2BJXoe4gU%2Fntqh9R%2FbV0oNUxWMvBQCIeeod2Jmwnxd4MWUzTxHiViRIwLMv2jggE8MIIBODAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUhq9%2FAsq75Mna8qYxs5JXcoUtEZ0wHwYDVR0jBBgwFoAUrqFzUYgpVxHKDKn0sQpuTrhLTQcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MDlKMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAKBggqhkjOPQQDAgNIADBFAiEAwyEJBoOM5Gl0CTWcuidzwgJCLEIFZdTO1gXAfJAjVAkCIAha4bOEWNvrHa%2FhgVUWMGwcK4BCAl5bLB2Bk%2Bwj2ptO)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 02 Nov 22 15:41 UTC