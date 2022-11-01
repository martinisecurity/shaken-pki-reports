# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ATI SHAKEN 731J

Tested At: 01 Nov 22 20:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 270 day(s)\
Subject: CN=ATI SHAKEN 731J, O=ATI, L=Phoenix, ST=AZ, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://certificates.peeringhub.io/731J/731J.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDCzCCArKgAwIBAgIQNyvUUal5R9gT%2F%2BFIQ2BJhDAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjA3MjkxNjQxNTZaFw0yMzA3MjkxNjQxNTZaMFQxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJBWjEQMA4GA1UEBwwHUGhvZW5peDEMMAoGA1UECgwDQVRJMRgwFgYDVQQDDA9BVEkgU0hBS0VOIDczMUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASl2XWD4%2FIjAfLLR7jmPtKWLMtDt5k0t%2BCzY0gHfiT3cbzYAutw1mgokOtKAiSt0RbTLUSvnu0K38F9zlprNh2io4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFH0dqeL3vPPVM1XgzSotBdU0l2rOMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENzMxSjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDRwAwRAIge24YdW%2Bsat%2BHlz4hNCajkVnrBCVBRX4hRLzcbKOpEwUCIBcL%2BVrHrJ5I0W3xDig6bdC%2FuWkmHtTKMtSp%2FZ46Zemy)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 01/11/2022 at 20:34:21