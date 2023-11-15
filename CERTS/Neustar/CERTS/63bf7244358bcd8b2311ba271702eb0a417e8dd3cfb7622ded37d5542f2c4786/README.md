# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar UAT Certified Caller ID SHAKEN CA-2

Tested At: 15 Nov 23 17:17 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 3247 day(s)\
Subject: CN=Neustar UAT Certified Caller ID SHAKEN CA-2, OU=www.ccid-uat.neustar, O=Neustar Information Services Inc, C=US\
Issuer: CN=Neustar UAT Certified Caller ID SHAKEN Root CA, OU=www.ccid-uat.neustar, O=Neustar Information Services Inc, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIDTjCCAvSgAwIBAgIUHog6XvAP%2Bpd56UtRobANOY2g%2BIswCgYIKoZIzj0EAwIwgZAxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEdMBsGA1UECwwUd3d3LmNjaWQtdWF0Lm5ldXN0YXIxNzA1BgNVBAMMLk5ldXN0YXIgVUFUIENlcnRpZmllZCBDYWxsZXIgSUQgU0hBS0VOIFJvb3QgQ0EwHhcNMjIxMDA1MTY1NzA1WhcNMzIxMDA1MTY1NzA1WjCBjTELMAkGA1UEBhMCVVMxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMR0wGwYDVQQLDBR3d3cuY2NpZC11YXQubmV1c3RhcjE0MDIGA1UEAwwrTmV1c3RhciBVQVQgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAjjpjbfq4K7V8cBrDseAw5FHadaG8l9FBf%2BXSJ4n4ALNV02IBJYETUv2bo0nhJDPCXgsl5OhOO%2Fg65UoUlFcaijggErMIIBJzAPBgNVHRMBAf8EBTADAQH%2FMB8GA1UdIwQYMBaAFEgo2zrklWRvA6p5WPWbN4vVSppJMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBqgYDVR0fBIGiMIGfMIGcoD6gPIY6aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLXN0Zy5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMB0GA1UdDgQWBBRkJ8xhVmXx7qQOdkqQZYdezKLWMTAOBgNVHQ8BAf8EBAMCAQYwCgYIKoZIzj0EAwIDSAAwRQIhAP8lMUpPRY6TAeh%2BPOGrSSmbAZTCtLQOLwlnyvnHw63BAiAVfezrESR82EHI9YQ6zT6lI9A3RbYF786eXe69Do6zjQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 17:17 UTC