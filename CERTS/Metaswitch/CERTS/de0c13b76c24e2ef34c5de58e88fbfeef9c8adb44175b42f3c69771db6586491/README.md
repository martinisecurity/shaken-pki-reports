# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Clearwave SHAKEN Cert 9915

Tested At: 21 Nov 23 16:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 202 day(s)\
Subject: CN=Clearwave SHAKEN Cert 9915, O=Clearwave Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/3f50200812ae74381c792fbff3f17ef55608e51d

[View certificate details](https://understandingwebpki.com/?cert=MIICXDCCAgKgAwIBAgIQSEulivW1RzMnFMf%2Bsvdm4DAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYxMTE2MDExNFoXDTI0MDYxMDE2MDExNFowVTELMAkGA1UEBhMCVVMxITAfBgNVBAoMGENsZWFyd2F2ZSBDb21tdW5pY2F0aW9uczEjMCEGA1UEAwwaQ2xlYXJ3YXZlIFNIQUtFTiBDZXJ0IDk5MTUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASyyY9024OYfDRCQnKThDHaXm1tU3M2Co9ldyLq6CotxfcCuMAEfQagCJP6u7EjA%2FjZL%2BOPBBLF94%2FjQ3QP8EEuo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ5OTE1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFH02vrr7Gc85LqhU2zvWBoRqF9VUMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQDtPV6bBkAF0gQOU%2F%2FTvn4W1xbSHp5PX8sUklJY63jK6gIgPgE6gCJNjBR4CgL6u7XtRELTqJq%2FihtqoMzLEgP6gWE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC