# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Valley Telephone Cooperative Inc / VTX1 SHAKEN 2159

Tested At: 26 Aug 24 17:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 187 day(s)\
Subject: CN=Valley Telephone Cooperative Inc / VTX1 SHAKEN 2159, O=Valley Telephone Cooperative Inc / VTX1, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/4cd5b91a833e4ccd9c7f4dabacdd26e85cddf71d

[View certificate details](https://x509.io/?cert=MIIChTCCAiqgAwIBAgIQYfiql8m%2FYM3WyidkoCUKSzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDMwMjE3NTIwN1oXDTI1MDMwMTE3NTIwN1owfTELMAkGA1UEBhMCVVMxMDAuBgNVBAoMJ1ZhbGxleSBUZWxlcGhvbmUgQ29vcGVyYXRpdmUgSW5jIC8gVlRYMTE8MDoGA1UEAwwzVmFsbGV5IFRlbGVwaG9uZSBDb29wZXJhdGl2ZSBJbmMgLyBWVFgxIFNIQUtFTiAyMTU5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEtecX6rrlHM751lvlw96Dk%2BJ2R5IxNd4KTG7U4aR3awsV%2FrXHyhw9hDoow1Qh7tXAS6oRbmPiVQEv3to4UrOM6qOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjE1OTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBSEjdLkPGjIbhTq0FcLVTh8cxd11DAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEAuV2PkIA14osyRReLYVvtGaiG8hXqpOS4lBs2FyL00YECIQDG1gZ4%2Fhga%2FzkFLnpqc4wAVnv%2Ba1TVupeAyqNIWsaXDQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:03 UTC