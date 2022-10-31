# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 9714

Tested At: 31 Oct 22 16:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 3 day(s)\
Subject: CN=SHAKEN 9714, OU=SHAKEN, O=Grid4 Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/51a37c7a-5af2-439d-94ce-677fa750ee2f/4ed1cdeae5e330e85aa1a08b0499c57e.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9jCCApygAwIBAgIQUjFEpPpyM497RXsBMMIqlTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYxODM4MTdaFw0yMjExMDIxODM4MTZaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRHcmlkNCBDb21tdW5pY2F0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gOTcxNDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABG0anv6PM6Y5r4DAEQabyodTxUnxt7oKYYpgBwCJCuWhRcpZcJqjpMg7KrbnJz4Y5VOHb%2BSr1BIqDlxVeyLT7i2jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUgmYKI%2BONhsrn6ljqR0%2B9wmQx08EwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEOTcxNDAKBggqhkjOPQQDAgNIADBFAiEAkheVrm4yc8KgatqI2WsQpZKfzfrvUQiWAEsj6camAfgCIG6a3YrG3IzMSJHjPj4%2Fj%2F7yue9kQxAKZosrWkTa1%2FKR)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31/10/2022 at 16:43:22