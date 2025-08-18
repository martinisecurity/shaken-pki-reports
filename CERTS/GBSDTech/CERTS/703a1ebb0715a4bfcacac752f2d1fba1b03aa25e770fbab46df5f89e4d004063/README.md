# STIR/SHAKEN CA Ecosystem Compliance

## Certificate FracTEL LLC SHAKEN

Tested At: 18 Aug 25 21:05 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 990 day(s)\
Subject: CN=FracTEL LLC SHAKEN, O=FracTEL LLC, ST=FL, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://ssc.fractel.net/ssc/fractelssc.pem

[View certificate details](https://x509.io/?cert=MIICuTCCAl%2BgAwIBAgICEVMwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNTA1MDUxNjQ4MTFaFw0yODA1MDQxNjQ4MTFaME0xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJGTDEUMBIGA1UECgwLRnJhY1RFTCBMTEMxGzAZBgNVBAMMEkZyYWNURUwgTExDIFNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEyrR9uftVX6yrFzOT8ZMhqwm8v6qi9DJ2Wi8xwwiLhEy8aWunZfdV%2BIIy0rheC0cqdyFBpal0YJy9LbUT6xH9GjggEaMIIBFjAWBggrBgEFBQcBGgQKMAigBhYEOTY1SDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBThFAMPbTwx7DeT9c48Cjk2HbfeATAfBgNVHSMEGDAWgBSnbEgJk6C%2FQeATYXHmA%2B01hx3ugTAOBgNVHQ8BAf8EBAMCB4AwgYQGA1UdHwR9MHsweaA%2BoDyGOmh0dHBzOi8vYXV0aGVudGljYXRlLWV4dC1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiN6Q1MDMxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEExEzARBgNVBAMMClNUSS1QQSBDUkwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMAoGCCqGSM49BAMCA0gAMEUCIQDX3z5DW%2BjkBzwBkMQewj8GfC6KQ3xvPqzJydsxYNq73QIgUJ4%2BlkT5La%2B9sCABwHUtzeDhGbudkGBI7DzDMZZ2r8Q%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 965H', but common name is 'FracTEL LLC SHAKEN' |


Generated: 18 Aug 25 21:13 UTC