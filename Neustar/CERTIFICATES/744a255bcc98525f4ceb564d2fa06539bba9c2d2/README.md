# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate 744a255bcc98525f4ceb564d2fa06539bba9c2d2
Tested At: 2022-10-28 18:22:34 +0000 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 540 day(s)\
Subject: C=US, ST=KS, L=Warmego, O=Warmego Telecom Company Inc, OU=Warmego Telecom, CN=Warmego\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1

Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/39.203

View: [Click to view](https://understandingwebpki.com/?cert=MIIECTCCAvGgAwIBAgIUIlarpqVOriw3m6tk0FNN1LEGW9EwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA0MTkyMzQ0NDRaFw0yNDA0MTkyMzQ0NDRaMH4xEDAOBgNVBAMMB1dhcm1lZ28xGDAWBgNVBAsMD1dhcm1lZ28gVGVsZWNvbTEkMCIGA1UECgwbV2FybWVnbyBUZWxlY29tIENvbXBhbnkgSW5jMRAwDgYDVQQHDAdXYXJtZWdvMQswCQYDVQQIDAJLUzELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARASjZTiZHIOFoWJ65zEp8vSwhSoi3JlA4%2BRIYznSOfslrq9JpRW752L45bxAnNAwE%2FKE9UydZKEvyb%2FRb4UUD8o4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBRBP7mLdJtjGlsmCYMBAUbyV9a1pDAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDE4NDUwDQYJKoZIhvcNAQELBQADggEBAJhzfMJmvTVjHATvRuJFWYWh7HU589pMyvBWTHhKoVQi6OiOnYScCetOXlK5QuJxMcIa15G1qtiLPUzIV6siby9MTNO9sscj9kb5KfN91b2utKnk7EJVcUrlLW27xYpKFhskt4iRs5YGDvQoC0EKtLnRHbnSQaki97SLZop%2Fkt3jDK3VxUQKXReS0HphhUTi%2FK9rc2WGLsYQ90f0%2FGuK7aXFJX9dNTD837%2BLCDaYfhvYGRWV1R%2FNBQ4zZRO1nKmwtP0J4Hidhk6z2%2Fy5cJgb6xGPZ1Ya8WP3bxzB79DM0T5e%2BSaLpAL9DFotW32l201gKfIiJqKede6Pl1bOlfku85E%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_sti_crl_distribution_not_reachable | error | ATIS-1000080 | Unable to retrieve CRL specified in CRLdp from allow listed IP address |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- w_cp_1_3_subject_email
- e_sti_serial_number
- e_sti_subject_cn
- e_sti_signature_algorithm
- e_sti_extension_unknown
- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 18:22:55