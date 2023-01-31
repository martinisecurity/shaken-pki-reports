# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Valley Telephone Cooperative Inc / VTX1 SHAKEN 2159

Tested At: 31 Jan 23 21:45 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 760 day(s)\
Subject: CN=Valley Telephone Cooperative Inc / VTX1 SHAKEN 2159, O=Valley Telephone Cooperative Inc / VTX1, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/6447136b7530812e9c9a848c3919f7947970fc7a

[View certificate details](https://understandingwebpki.com/?cert=MIIChTCCAiqgAwIBAgIQYfiql8m%2FYM3WyidkoCUKSzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDMwMjE3NTIwN1oXDTI1MDMwMTE3NTIwN1owfTELMAkGA1UEBhMCVVMxMDAuBgNVBAoMJ1ZhbGxleSBUZWxlcGhvbmUgQ29vcGVyYXRpdmUgSW5jIC8gVlRYMTE8MDoGA1UEAwwzVmFsbGV5IFRlbGVwaG9uZSBDb29wZXJhdGl2ZSBJbmMgLyBWVFgxIFNIQUtFTiAyMTU5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEtecX6rrlHM751lvlw96Dk%2BJ2R5IxNd4KTG7U4aR3awsV%2FrXHyhw9hDoow1Qh7tXAS6oRbmPiVQEv3to4UrOM6qOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjE1OTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBSEjdLkPGjIbhTq0FcLVTh8cxd11DAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEAuV2PkIA14osyRReLYVvtGaiG8hXqpOS4lBs2FyL00YECIQDG1gZ4%2Fhga%2FzkFLnpqc4wAVnv%2Ba1TVupeAyqNIWsaXDQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


Generated: 31 Jan 23 21:50 UTC