# STIR/SHAKEN CA Ecosystem Compliance

## Certificate 846B

Tested At: 31 Oct 22 20:40 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 566 day(s)\
Subject: C=US, ST=Texas, L=Hallsville, O=Tim Ron Enterprises\\, LLC, OU=SHAKEN, CN=846B\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://prod001-cr.rbbnidhub.com/OaPccQ6Mgz/sign-cert2

View: [Click to view](https://understandingwebpki.com/?cert=MIIEADCCAuigAwIBAgIUAak5C9kCINIjZPYQa8fE6klP19MwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MTkxNDAyMTFaFw0yNDA1MTkxNDAyMTFaMHUxDTALBgNVBAMMBDg0NkIxDzANBgNVBAsMBlNIQUtFTjEhMB8GA1UECgwYVGltIFJvbiBFbnRlcnByaXNlcywgTExDMRMwEQYDVQQHDApIYWxsc3ZpbGxlMQ4wDAYDVQQIDAVUZXhhczELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARjf3%2BD8yyzm0SnlD2H51u1ayapMLEmNfvfshD8GdMtOJNXZyrL3qJX4qmYaTG4%2FOFiUBw5kcWmDaDXJ47O5VPDo4IBSDCCAUQwFgYIKwYBBQUHARoECjAIoAYWBDg0NkIwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBSJa8uGyv8%2BhcrgmKQgUK9S%2BOIqxDAOBgNVHQ8BAf8EBAMCB4AwDQYJKoZIhvcNAQELBQADggEBAJ5GIQR5miE2aC5M5lZAkmc%2FsGxGe3tQY8K9Oek4%2FW2X5iRKP3tO97iVZsA%2BBf%2B4eW0XxTo6GME45acLUW0lL8PEK9HvBpWFjjcAKkeyew2qkivsZNk2%2FldjqXWKpYn%2FCoCh64FNjC6IR0KHdwmdmt8HiJyK6cRmVHwVil25crRswth%2BZDNMeGMeLu69yX1olCcS9i%2FLMjQDcV6M6QiJQbH%2FVHn754L591lP6tuvzWqYCgVziTxNMjdcaB6EB4ip6OdICaO2T1VWCjpGvSqKGL6DmkcCbMUNiCzbnjbYZp1%2Fg0ctyyJZ1acv0Pul9tyrcLkwK6zmci9wtV8HcfeTZ2M%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 20:47:45