# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar Certified Caller ID SHAKEN CA-2

Tested At: 30 Jan 23 23:10 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 3536 day(s)\
Subject: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN Root CA, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIDOTCCAuCgAwIBAgIUWrGUQ%2BfMvcWusHLkptixbxD9NtswCgYIKoZIzj0EAwIwgYgxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEzMDEGA1UEAwwqTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBSb290IENBMB4XDTIyMTAwNTE3MjY0NFoXDTMyMTAwNTE3MjY0NFowgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEgmiNQIrfTzKNW92OodClrQUs7HoCpT4TtPeVm1r%2Ff%2FGwkuRUR1OytazCx8ZzsheI99U07qvCbf516MaRiBuJ06OCAScwggEjMA8GA1UdEwEB%2FwQFMAMBAf8wHwYDVR0jBBgwFoAUFNWx4g9D2wES8o4AOlMKXL1z19AwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAdBgNVHQ4EFgQUgk4V%2F%2F6famdR5MiXx210w%2FxlRXgwDgYDVR0PAQH%2FBAQDAgEGMAoGCCqGSM49BAMCA0cAMEQCID4c2oB7s%2BofBkxcF%2BSNM4NE%2Fe7yqzpEERonOPvehQNvAiBjE2mjr3eb2eDGPcnHWxVL7K7hACMI%2FMtfGMYdyXvNYQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 30 Jan 23 23:10 UTC