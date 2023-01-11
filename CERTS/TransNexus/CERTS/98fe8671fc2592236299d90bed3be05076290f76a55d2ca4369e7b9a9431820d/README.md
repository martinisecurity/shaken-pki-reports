# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 721J

Tested At: 11 Jan 23 20:37 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -91 day(s)\
Subject: CN=SHAKEN 721J, OU=SHAKEN, O=True IP Solutions, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/a1df4bf4-8858-47bb-9388-835c5c7cb5c4/c444a025b79edfc6845719a4946375dd.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApmgAwIBAgIQSuf6PB%2FuBptTk2zpxTQCNzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDUyMDIwMTBaFw0yMjEwMTIyMDIwMDlaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFUcnVlIElQIFNvbHV0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzIxSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABH1JQi2MSxvoPv34ll80EDc%2B6OO7uSwtQBOGsPOXijEtL405UEyoHVAouzQc0rMmWw4fEwT2wX6QkbjlBUuyoCajggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUWtGIhoFzLHl5zkz0lZMsUyMNOfowHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzIxSjAKBggqhkjOPQQDAgNIADBFAiBtml4RHIEs5y3A6oCxwq8dxJfN0rcLKBqpxyloBtd5KQIhAOjubK13W9rKtTIXHXzBGa97Vjp%2F487rvlPgWDxpjvKG)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 11 Jan 23 21:04 UTC