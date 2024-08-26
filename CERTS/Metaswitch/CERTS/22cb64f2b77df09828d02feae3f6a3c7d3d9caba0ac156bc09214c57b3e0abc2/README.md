# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Allstream SHAKEN 4130

Tested At: 26 Aug 24 17:43 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 944 day(s)\
Subject: CN=Allstream SHAKEN 4130, OU=QCall, O=Allstream Business US\\, LLC, L=Vancouver, ST=WA, C=US\
Issuer: O=Metaswitch Networks, C=GB, CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cert-stir-us.allstream.com/certs/allstreamcertchain.crt

[View certificate details](https://x509.io/?cert=MIICtjCCAlygAwIBAgIQHUDcNkbSOgfM4BJB%2Bv9zFzAKBggqhkjOPQQDAjBYMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMQswCQYDVQQGEwJHQjEcMBoGA1UECgwTTWV0YXN3aXRjaCBOZXR3b3JrczAeFw0yNDAzMjgxNjQzMzNaFw0yNzAzMjgxNjQzMzNaMIGDMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEjAQBgNVBAcMCVZhbmNvdXZlcjEjMCEGA1UECgwaQWxsc3RyZWFtIEJ1c2luZXNzIFVTLCBMTEMxDjAMBgNVBAsMBVFDYWxsMR4wHAYDVQQDDBVBbGxzdHJlYW0gU0hBS0VOIDQxMzAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATtIbEdLh4pzEpphtpHXNCopXerARrFZNjqA6%2FJ2H6SJlnL8W4SM3bejqozCDNWyhCHCk6yZ5ECi9DwmvyaDgWbo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgbAMBYGCCsGAQUFBwEaBAowCKAGFgQ0MTMwMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwHQYDVR0OBBYEFOaP8FBgZL1hpXc4pa%2Fx5YQJZRnQMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCBu99J1iVdcbZzBmzlwwM%2Bg9q3R11BjFOaurEIEkWGIAIgaTQdelmxsOF8kRoF%2FWY2is0XOMjLSirhGS%2F01i%2Fihnc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 26 Aug 24 18:03 UTC