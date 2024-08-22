# STIR/SHAKEN CA Ecosystem Compliance

## Certificate FracTEL LLC SHAKEN

Tested At: 22 Aug 24 15:59 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 274 day(s)\
Subject: CN=FracTEL LLC SHAKEN, O=FracTEL LLC, ST=Florida, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://ssc.fractel.net/ssc/fractelssc.pem

[View certificate details](https://x509.io/?cert=MIICnTCCAkSgAwIBAgICERMwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNDA1MjMxNjUxMjNaFw0yNTA1MjIxNjUxMjNaMFIxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMRQwEgYDVQQKDAtGcmFjVEVMIExMQzEbMBkGA1UEAwwSRnJhY1RFTCBMTEMgU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE4vvK0G%2FZ%2Bu296Pqdh2sRDD8G3Ar052xlqDRgwxfFPR76Uk7FY01D%2Boa0THk5t3%2BmoPHN7L2KYkC%2BYZvO3sVpa6OB%2BzCB%2BDAWBggrBgEFBQcBGgQKMAigBhYEOTY1SDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQgruypwj0N5JlQe1Ad8fAQXJyXbjAfBgNVHSMEGDAWgBR2xOfJOx4nEedS1E2tePaQBz6blDAOBgNVHQ8BAf8EBAMCB4AwZwYDVR0fBGAwXjBcoCGgH4YdaHR0cHM6Ly9zdGktcGEuY29tL3NoYWtlbi9jcmyiN6Q1MDMxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEExEzARBgNVBAMMClNUSS1QQSBDUkwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMAoGCCqGSM49BAMCA0cAMEQCIDaacfbhqiU5jiFosPp8c%2BUb%2B0KL4zsE4j47X4hNhO4wAiB4lkf51ah5h4C406%2B96XUVTBA039%2F2eCi3rqq6cj%2BKcg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 965H', but common name is 'FracTEL LLC SHAKEN' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |


Generated: 22 Aug 24 16:06 UTC