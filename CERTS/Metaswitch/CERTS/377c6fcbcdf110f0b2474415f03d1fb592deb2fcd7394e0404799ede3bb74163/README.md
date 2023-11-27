# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Chariton Valley SHAKEN 250A

Tested At: 27 Nov 23 23:10 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 887 day(s)\
Subject: CN=Chariton Valley SHAKEN 250A, O=Chariton Valley Communications Corporation, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/0ba2e829dff0ca7fef842872ab82485c69f76281

[View certificate details](https://understandingwebpki.com/?cert=MIICbjCCAhWgAwIBAgIQKXhZnVTHqxfjd2l8L9ddYDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUwMzE2NTU0N1oXDTI2MDUwMjE2NTU0N1owaDELMAkGA1UEBhMCVVMxMzAxBgNVBAoMKkNoYXJpdG9uIFZhbGxleSBDb21tdW5pY2F0aW9ucyBDb3Jwb3JhdGlvbjEkMCIGA1UEAwwbQ2hhcml0b24gVmFsbGV5IFNIQUtFTiAyNTBBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEu3IBjT%2Fd5opSaW6HhrXNEeKXePCHyMMYF8Zc9DWGXv3qTKI9eiS%2B2MSx4diRFG38ElzL%2FM16mjAWzDfalJeBC6OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjUwQTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBSdnuNic771GlKqwhZlYEqJC5rgwTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiAneB1%2BOOdCQIFScs0r6%2FQL6Jik2H7uaK5Y2PeW2jA81QIgFvUXRp2weed%2F08VFmAgA3hCNeVuVs9kbjZKQZJhO39Q%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 27 Nov 23 23:28 UTC