# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 9714

Tested At: 05 Jan 23 18:06 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -78 day(s)\
Subject: CN=SHAKEN 9714, OU=SHAKEN, O=Grid4 Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/51a37c7a-5af2-439d-94ce-677fa750ee2f/859e58181583d0d83e3f5f21516a6cff.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9jCCApygAwIBAgIQa6vX1%2F%2BpJAAaobRM7mYJXTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTExODM2NThaFw0yMjEwMTgxODM2NTdaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRHcmlkNCBDb21tdW5pY2F0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gOTcxNDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJqKiHuBGIKqEZ%2BZfQmjlTuN1B80gz3NO%2FfQJWR1%2FBP6BDzzPUZGDYS5PtrM1aR9jQsp1V46yeB2Axbqwfzyr3ujggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQULwuz%2FXLLUMxc3y1IyZo2d46EDhMwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEOTcxNDAKBggqhkjOPQQDAgNIADBFAiEAgnPBCQmlUmzGK0fPFHuqkhBJozikcU2rcecZotNCDe8CICBq8DXn3LIAOFjW1oJBdR9YV9GF8qlGR5nwPBVmzi10)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 05 Jan 23 18:35 UTC