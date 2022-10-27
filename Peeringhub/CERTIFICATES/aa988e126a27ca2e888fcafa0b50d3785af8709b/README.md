# STIR/SHAKEN CA Ecosystem Compliance
## Peeringhub

### Certificate aa988e126a27ca2e888fcafa0b50d3785af8709b
Tested At: 2022-10-27 18:56:40 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 275 day(s)\
Subject: CN=ATI SHAKEN 731J, O=ATI, L=Phoenix, ST=AZ, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US

Link: https://certificates.peeringhub.io/731J/731J.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDCzCCArKgAwIBAgIQNyvUUal5R9gT%2F%2BFIQ2BJhDAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjA3MjkxNjQxNTZaFw0yMzA3MjkxNjQxNTZaMFQxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJBWjEQMA4GA1UEBwwHUGhvZW5peDEMMAoGA1UECgwDQVRJMRgwFgYDVQQDDA9BVEkgU0hBS0VOIDczMUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASl2XWD4%2FIjAfLLR7jmPtKWLMtDt5k0t%2BCzY0gHfiT3cbzYAutw1mgokOtKAiSt0RbTLUSvnu0K38F9zlprNh2io4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFH0dqeL3vPPVM1XgzSotBdU0l2rOMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENzMxSjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDRwAwRAIge24YdW%2Bsat%2BHlz4hNCajkVnrBCVBRX4hRLzcbKOpEwUCIBcL%2BVrHrJ5I0W3xDig6bdC%2FuWkmHtTKMtSp%2FZ46Zemy)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 27/10/2022 at 18:57:26