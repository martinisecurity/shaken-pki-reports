# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate 51cef5d8e5387e2cf7c8cb171737e2cb4216a58c
Tested At: 2022-10-27 22:32:02 +0000 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 554 day(s)\
Subject: C=US, ST=AR, L=LittleRock, O=WindstreamServices, OU=WindstreamCommunication, CN=WindstreamCommunication\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1

Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11162.10187

View: [Click to view](https://understandingwebpki.com/?cert=MIIEHDCCAwSgAwIBAgIUYJc9NrwJpQ0WGUL4zAUy2UhOuRswDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MDMyMTExNThaFw0yNDA1MDMyMTExNThaMIGQMSAwHgYDVQQDDBdXaW5kc3RyZWFtQ29tbXVuaWNhdGlvbjEgMB4GA1UECwwXV2luZHN0cmVhbUNvbW11bmljYXRpb24xGzAZBgNVBAoMEldpbmRzdHJlYW1TZXJ2aWNlczETMBEGA1UEBwwKTGl0dGxlUm9jazELMAkGA1UECAwCQVIxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEwzFU6EQ6xAzXwVrVjgk9aPypuQ%2FQPz8Tsd7rEDeiOWUBsiHkbZDoZbqxE70Sp2G8OMnzLqEmgDFsuTnGEhw0MKOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQU2CeKIfWvomorUpz9czuGpEjkDy0wDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ0MTUxMA0GCSqGSIb3DQEBCwUAA4IBAQBG6aDv9UZewqbbBKjTP94XKbuTnyj0i773qLr%2Bl4UvBsyhqbaeafVDl9vNriXSeGPX%2FP%2F7xhoGx6jxm9FdrgteiYarK7EqVjJCbVSCpFflc4f6iA6vqXkc9zSPU2O7skL%2Bv79QkH3b%2B49idMM8De%2FwwxstaxpsJbdNDDdhJLU%2FL8187FHY5qrNuRUWuYs2qe4zCDbgqGMeIQ3ZlSj4ZquYKFk%2FMcOh73maNYX4kQ3DQQ0qt0IbSib9WS6eAySEGIeEExx5eXcjjJcQrC49rVLQbwNP1AN4rPizXPfDqIFFhf6%2F2m3Qxd3Y%2FZIFC1%2FQUKPMlkeGH87PG6TNXJxA6i6F)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_sti_crl_distribution_not_reachable | error | ATIS-1000080 | Unable to retrieve CRL specified in CRLdp from allow listed IP address |

### Not Effective

- e_sti_extension_unknown
- e_sti_subject_cn
- e_sti_signature_algorithm
- w_cp_1_3_subject_email
- w_cp1_3_subject_rdn_unknown
- e_sti_serial_number
- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 22:33:03