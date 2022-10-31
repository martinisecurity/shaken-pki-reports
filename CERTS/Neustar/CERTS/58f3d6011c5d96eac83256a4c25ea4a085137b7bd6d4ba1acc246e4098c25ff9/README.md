# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Shaken

Tested At: 31 Oct 22 16:41 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 479 day(s)\
Subject: C=US, ST=CO, L=Denver, O=Teliax Inc., OU=Teliax Inc., CN=Shaken\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11316.10181.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIID8zCCAtugAwIBAgIUK4qW235HUJnS6WWxgWbbp7WMVT8wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAyMjExMzU3MjJaFw0yNDAyMjIxMzU3MjJaMGgxDzANBgNVBAMMBlNoYWtlbjEUMBIGA1UECwwLVGVsaWF4IEluYy4xFDASBgNVBAoMC1RlbGlheCBJbmMuMQ8wDQYDVQQHDAZEZW52ZXIxCzAJBgNVBAgMAkNPMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPM%2FHRBpOoitBJ3Hv7hftxio6lUt4b2pQr4jhhkE5a5PSDFQ1rpMgcG4wruy8MtjGgC%2BgFPcMuAFyd7AX5CdosijggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFFY61eI5bVULQksViry9C7Iibq1rMA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYENTY3RzANBgkqhkiG9w0BAQsFAAOCAQEAcOGoW6XU6r7tnEm%2BY%2BCO94%2BctvaMg%2F%2FUdaN%2BTQ9iLjluVtm33p5TlxszK6jvr17RQXHdJ4RqnOTWNWXJ4s29Ix2xPMTC%2FsieEHL%2B2jOoMEyY4YCxV3ZyHfVd%2FB3tLQgAARF1i02km%2BIWDLKt2IUR5h8U%2Bls4lq2u%2FsQVODsO83tjt7W4QYBtB9UYHVPCMKAygBt9r9LjP%2Fiy3loGxaDr5dB4Pgo9Lomggm1MSmAsak4oh8z7NNLiDPZdI7AncwUmLKQJBMlgG1UOFGVNGMTxR0QrEOcpComv%2FSRyM2LnJvJngAZEbee8AZyQLIA1X0aFEVsjDNEt6SmnMVNkS0RhlA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_sti_crl_distribution_not_reachable](../../ISSUES/e_sti_crl_distribution_not_reachable/README.md) | error | ATIS&#x2011;1000080 | Unable to retrieve CRL specified in CRLdp from allow listed IP address |
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


Generated: 31/10/2022 at 16:43:22