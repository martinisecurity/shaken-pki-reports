# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Twin Valley SHAKEN 1840

Tested At: 18 Aug 25 20:05 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 1019 day(s)\
Subject: CN=Twin Valley SHAKEN 1840, O=Twin Valley, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://cdn-cr.cgah.tnsi.com/certs/f292d2a556a66a810fed9e79097db082bd9a66a2

[View certificate details](https://x509.io/?cert=MIICsjCCAlegAwIBAgICEXEwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNTA2MDMxNDQyNTBaFw0yODA2MDIxNDQyNTBaMEUxCzAJBgNVBAYTAlVTMRQwEgYDVQQKDAtUd2luIFZhbGxleTEgMB4GA1UEAwwXVHdpbiBWYWxsZXkgU0hBS0VOIDE4NDAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASqVYixPWOdt3LduVyMmY7beNmCd6Kn4HnV2pHl6eQUBBtct%2F5pY9rL2C2ZGizy%2Bsziq5kbKn%2BlclCN03y4tsQOo4IBGjCCARYwFgYIKwYBBQUHARoECjAIoAYWBDE4MzMwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUQy%2Bnj%2F03UMzfhHNtUOgv4%2BoOhLQwHwYDVR0jBBgwFoAUp2xICZOgv0HgE2Fx5gPtNYcd7oEwDgYDVR0PAQH%2FBAQDAgeAMIGEBgNVHR8EfTB7MHmgPqA8hjpodHRwczovL2F1dGhlbnRpY2F0ZS1leHQtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsojekNTAzMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMRMwEQYDVQQDDApTVEktUEEgQ1JMMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAKBggqhkjOPQQDAgNJADBGAiEApcZgGuZbr9zxHIWnNPzBcnYhLIGGTWitI%2BpNssWD6WgCIQDONpe7dmLthhtZu1rqNCOhmDMu%2FzvyTdI0w9p%2BJIn9VA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1833', but common name is 'Twin Valley SHAKEN 1840' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |


Generated: 18 Aug 25 21:13 UTC