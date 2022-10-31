# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 721J

Tested At: 31 Oct 22 19:20 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 3 day(s)\
Subject: CN=SHAKEN 721J, OU=SHAKEN, O=True IP Solutions, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/a1df4bf4-8858-47bb-9388-835c5c7cb5c4/4e583b9d11e190ea91c0687b930b92db.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8zCCApmgAwIBAgIQQ3nsBXr4zZJmVYstriKquzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMDIyNDlaFw0yMjExMDIyMDIyNDhaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFUcnVlIElQIFNvbHV0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzIxSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKs14BofyKCvgTSzaG21vxyMFCJce0U1jKeimCT%2Bl2LGKFXjKN0Ph%2FF4%2FfEOLkZ84VOnLk0ch3eFmU6oMyfraVWjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUHSpSEGQ1r4oxav3nJfRU8NqGtngwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzIxSjAKBggqhkjOPQQDAgNIADBFAiAF%2BZIfAWsKc20pk3x1kPFYa4zXeXwpQsz3IayL1qfI1AIhAMQrYf29Qp6yvs0NKTFhhRxfoKOwH8cTFHHuW%2FR2%2FvKq)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31/10/2022 at 19:21:49