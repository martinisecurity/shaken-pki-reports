# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN-651J

Tested At: 18 Aug 25 21:08 UTC\
Initial Validity Period: 398 day(s)\
Remaining Validity Period: 202 day(s)\
Subject: CN=SHAKEN-651J, O=Tata Communications Canada Ltd, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN CA-2, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Link: https://sticr-cstga.ccid.neustar/api/v1/certificate/9d7d837bf5361ef723537b554afa276a.pem

[View certificate details](https://x509.io/?cert=MIIDVjCCAvugAwIBAgIUFyfzY1Ptm03cLMQgt0rfiTa8iswwCgYIKoZIzj0EAwIwgY8xCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTd3d3LmNhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMjAeFw0yNTAyMDMxODQxMTZaFw0yNjAzMDgxODQxMTZaMEwxCzAJBgNVBAYTAkNBMScwJQYDVQQKDB5UYXRhIENvbW11bmljYXRpb25zIENhbmFkYSBMdGQxFDASBgNVBAMMC1NIQUtFTi02NTFKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE35MRFXqdyrCChI%2BKazbqdoU74HQa0wuUvnv5FBfecKliD8hiLSj%2Bid1cLvMmuGhdDDwagC8tjxw0E4rE3ucRbKOCAXUwggFxMBYGCCsGAQUFBwEaBAowCKAGFgQ2NTFKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAULR03ZEhPDh1sSKFk2FnVptAGwmEwFgYDVR0gBA8wDTALBgkrBgEEAYO3HwEwgeAGA1UdHwSB2DCB1TCB0qAzoDGGL2h0dHBzOi8vc3RpcGEtY3JsLWNzdGdhLmNjaWQubmV1c3Rhci9hcGkvdjEvY3JsooGapIGXMIGUMQswCQYDVQQGEwJDQTEpMCcGA1UECgwgTmV1c3RhciBJbmZvcm1hdGlvbiBTZXJ2aWNlcyBJbmMxITAfBgNVBAsMGHN0aXBhLWNzdGdhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBTVEktUEEgUm9vdCBDQTAdBgNVHQ4EFgQUhSQyoJ9jC2u0Hh0SEXRaDa4cOjcwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCPlIiqc9ZiKcQbv%2BOmgUj0EtBVB82I0%2Btpd4fVxiDYOAIhAPZDKgyUHXvyZQfLzzsmTlGIFDBnG3Or07u4Z38HnCa4)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 651J', but common name is 'SHAKEN-651J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 1.3.6.1.4.1.56223.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 18 Aug 25 21:13 UTC