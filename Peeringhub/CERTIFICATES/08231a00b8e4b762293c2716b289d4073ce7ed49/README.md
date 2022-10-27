# STIR/SHAKEN CA Ecosystem Compliance
## Peeringhub

### Certificate 08231a00b8e4b762293c2716b289d4073ce7ed49
Tested At: 2022-10-27 22:42:55 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 277 day(s)\
Subject: CN=Voiceterm SHAKEN 240K, O=Voiceterm, L=Phoenix, ST=AZ, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US

Link: https://certificates.peeringhub.io/240K/240K.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDGTCCAr6gAwIBAgIQOJ%2F%2Bvu5482JqSaG%2BT0jweDAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjA3MzEwOTA5MTZaFw0yMzA3MzEwOTA5MTZaMGAxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJBWjEQMA4GA1UEBwwHUGhvZW5peDESMBAGA1UECgwJVm9pY2V0ZXJtMR4wHAYDVQQDDBVWb2ljZXRlcm0gU0hBS0VOIDI0MEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASU2eh%2BOQkkEJw9Z8cTGvGS8Id5%2Bv3RCCawsHTmBi7%2BIEaptmqoCDkMQ482M%2BurHdj2NrS6oQQJe6eJARKyo39ao4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFA8oB3xJXX6%2BIQ3GL7WeCYkGaQcuMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYEMjQwSzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDSQAwRgIhANTVamvsVVNNGQKoTbZ%2BQOqMRFei2FZYJv2dZEFJutfpAiEA4NidJ98DqmpKBLhtV3DZK5DMjGykAmfSVUyKm052ENY%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 27/10/2022 at 22:44:50