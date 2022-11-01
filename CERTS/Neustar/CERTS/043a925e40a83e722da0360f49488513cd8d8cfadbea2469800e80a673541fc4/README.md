# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Ringcentral-ProdKeystore

Tested At: 01 Nov 22 22:11 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 466 day(s)\
Subject: C=US, ST=CA, L=Belmont, O=Ringcentral, OU=Ringcentral, CN=Ringcentral-ProdKeystore\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11155.10177

View: [Click to view](https://understandingwebpki.com/?cert=MIIEBjCCAu6gAwIBAgIUGTiH6BoLDybPxcL0pc9sqRDhj30wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAyMDkyMTQ5MThaFw0yNDAyMTAyMTQ5MThaMHsxITAfBgNVBAMMGFJpbmdjZW50cmFsLVByb2RLZXlzdG9yZTEUMBIGA1UECwwLUmluZ2NlbnRyYWwxFDASBgNVBAoMC1JpbmdjZW50cmFsMRAwDgYDVQQHDAdCZWxtb250MQswCQYDVQQIDAJDQTELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASAuzQlaVAp8lyoPH%2FSmoSMTklam%2F6BA50GUCl6FrClRR9TTlqrDdrpyUK5MfSPlnqb%2Bs%2Bp0RLlZ66reLVAr2Qvo4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBQuVq2vYpoiK01HBdcAUKdQh6kh6zAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDg4NkcwDQYJKoZIhvcNAQELBQADggEBAGH6R0tpApFYXPh2bYa%2B37f6n7gkWJ%2Fwd7iPJYFyyE%2FyvhvFhZW7tdDIjuyW1LcOwvltiG3k%2BtmwFHDcmn0LVVqK39RQVGjnjZjGjDkyqxaomqG6KNeQVhtF0DFVspttBSfOKwb4nTodhzMl88qnknIuqx%2Franz0IpGi0q%2F0k5l5Ba2p9OgGEpw%2FkSKKKLJFyad3oO4CTXuRMsv4P9YWLkwF38p59rfmYr%2BPwCIe12jbG%2FU6TwP4iiFtFopv1WuqLHFfomDWP2UpOhZR%2Fx3JXQTs9QBwdA7TSfvWJ%2B4034nsYTGU2JDmcVgtFqL%2FXLmTp7Ah9YHxuRj0eMDcUJLlevQ%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 22:19:34