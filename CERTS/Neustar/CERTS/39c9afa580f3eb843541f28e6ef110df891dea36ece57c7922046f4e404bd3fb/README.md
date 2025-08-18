# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN8304

Tested At: 18 Aug 25 20:05 UTC\
Initial Validity Period: 398 day(s)\
Remaining Validity Period: 167 day(s)\
Subject: CN=SHAKEN8304, OU=VOIP, O=Allstream Business Inc., L=Mississauga, ST=ON, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN CA-2, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Link: https://cert-stir-ca.allstream.com/certs/allstreamcertchain.crt

[View certificate details](https://x509.io/?cert=MIIDgDCCAyWgAwIBAgIUYEeSFcL1v6amAHOX9HCau9zwXhIwCgYIKoZIzj0EAwIwgY8xCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTd3d3LmNhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMjAeFw0yNDEyMzAxODMzNThaFw0yNjAyMDExODMzNThaMHYxCzAJBgNVBAYTAkNBMQswCQYDVQQIDAJPTjEUMBIGA1UEBwwLTWlzc2lzc2F1Z2ExIDAeBgNVBAoMF0FsbHN0cmVhbSBCdXNpbmVzcyBJbmMuMQ0wCwYDVQQLDARWT0lQMRMwEQYDVQQDDApTSEFLRU44MzA0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESTNqQOjVqswoXb%2FDP4TH2iTeYSeTJcLuzVOr%2FUAm7sjc7%2FZOblUFNgoar7VCtOcpVJ1Atr%2B14wXz2SldzBZ0OaOCAXUwggFxMBYGCCsGAQUFBwEaBAowCKAGFgQ4MzA0MAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAULR03ZEhPDh1sSKFk2FnVptAGwmEwFgYDVR0gBA8wDTALBgkrBgEEAYO3HwEwgeAGA1UdHwSB2DCB1TCB0qAzoDGGL2h0dHBzOi8vc3RpcGEtY3JsLWNzdGdhLmNjaWQubmV1c3Rhci9hcGkvdjEvY3JsooGapIGXMIGUMQswCQYDVQQGEwJDQTEpMCcGA1UECgwgTmV1c3RhciBJbmZvcm1hdGlvbiBTZXJ2aWNlcyBJbmMxITAfBgNVBAsMGHN0aXBhLWNzdGdhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBTVEktUEEgUm9vdCBDQTAdBgNVHQ4EFgQUs%2BcNp3MdvlCH1MYviyfThZghpCYwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDcKWiHMraFExPNw%2FAx1nNaFJWy29AxirDEl31pkih6fQIhAPKpIBP%2FdSg4%2B2F7h1T5MUZq%2FLihrdH7KMC9K6vNyR1R)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'SHAKEN8304' does not contain 'SHAKEN' |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 8304', but common name is 'SHAKEN8304' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 1.3.6.1.4.1.56223.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 18 Aug 25 21:13 UTC