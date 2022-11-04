# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Call48 SHAKEN Cert 505J

Tested At: 04 Nov 22 01:11 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 430 day(s)\
Subject: CN=Call48 SHAKEN Cert 505J, O=Call48, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/27619989593575ff99cddfc6b207ba73fed5de1e

[View certificate details](https://understandingwebpki.com/?cert=MIICRzCCAe2gAwIBAgIQeAW6uFjvX%2Bmzp%2BwrVaMWCDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDEwNzE2MTcwOFoXDTI0MDEwNzE2MTcwOFowQDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBkNhbGw0ODEgMB4GA1UEAwwXQ2FsbDQ4IFNIQUtFTiBDZXJ0IDUwNUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATssH1beXS%2BZQrXa1EE8SEwwJu3GRb2ThiUFZ9xyB%2Bh%2B5qQMmW%2BkFEHaXP%2BIAFwpisw%2Bgloa2N52pyz%2BC%2FCg9TUo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ1MDVKMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFMRJ9BaJzMlVK4auDorqr8MenmV8MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIC5Ng2VejjymxXhs%2FeM08xe2v0rcnPj9vFLCKuh6CBGNAiEAo1pPvMlV%2ByAcrl%2B7A3R2lnpTp%2BjdqoACut2mBQo5CFg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Nov 22 01:11 UTC