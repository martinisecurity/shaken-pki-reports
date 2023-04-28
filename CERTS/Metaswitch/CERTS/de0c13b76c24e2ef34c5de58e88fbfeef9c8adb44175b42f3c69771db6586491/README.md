# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Clearwave SHAKEN Cert 9915

Tested At: 28 Apr 23 02:03 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 410 day(s)\
Subject: CN=Clearwave SHAKEN Cert 9915, O=Clearwave Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/f63503c138e2b55933317f09cc5a1cc4e8a3de1d

[View certificate details](https://understandingwebpki.com/?cert=MIICXDCCAgKgAwIBAgIQSEulivW1RzMnFMf%2Bsvdm4DAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYxMTE2MDExNFoXDTI0MDYxMDE2MDExNFowVTELMAkGA1UEBhMCVVMxITAfBgNVBAoMGENsZWFyd2F2ZSBDb21tdW5pY2F0aW9uczEjMCEGA1UEAwwaQ2xlYXJ3YXZlIFNIQUtFTiBDZXJ0IDk5MTUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASyyY9024OYfDRCQnKThDHaXm1tU3M2Co9ldyLq6CotxfcCuMAEfQagCJP6u7EjA%2FjZL%2BOPBBLF94%2FjQ3QP8EEuo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ5OTE1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFH02vrr7Gc85LqhU2zvWBoRqF9VUMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQDtPV6bBkAF0gQOU%2F%2FTvn4W1xbSHp5PX8sUklJY63jK6gIgPgE6gCJNjBR4CgL6u7XtRELTqJq%2FihtqoMzLEgP6gWE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Apr 23 02:17 UTC