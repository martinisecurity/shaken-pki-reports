# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate cb1430446ee425f297b477264f245934a792e826
Tested At: 2022-10-28 18:22:46 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 3 day(s)\
Subject: CN=SHAKEN 9714, OU=SHAKEN, O=Grid4 Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/51a37c7a-5af2-439d-94ce-677fa750ee2f/2574096f567842a90e74c70e07de3333.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9zCCApygAwIBAgIQXNSiM4JXk6Ykh3xe%2BkHEVTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMxODM3NTZaFw0yMjEwMzAxODM3NTVaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRHcmlkNCBDb21tdW5pY2F0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gOTcxNDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPTfDODBKIbXr1zlyexjVSLGyCZngYC8jKsL1MhAwXSDZPXK1JldD5r7eUPTDBy%2ByueHMykAXl6OtYY0e0L86cWjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUhp9s0LnXBl%2BaW8xk7y0c%2B%2F984NwwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEOTcxNDAKBggqhkjOPQQDAgNJADBGAiEAyBiCWvhwYRzwtWD3rdsuHYEpnhHAidL84k0uXJAKY%2FsCIQCJD1afkKARqQ6TTcjpNEOHRrEK3hbwI0P6N61HQNegHg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 18:22:55