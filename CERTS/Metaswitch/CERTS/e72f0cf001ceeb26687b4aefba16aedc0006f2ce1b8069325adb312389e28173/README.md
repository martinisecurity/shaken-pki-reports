# STIR/SHAKEN CA Ecosystem Compliance

## Certificate DTC Communications SHAKEN Cert 0562

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 226 day(s)\
Subject: CN=DTC Communications SHAKEN Cert 0562, O=DTC Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/e5db82c70d06ef1ac7577b261ae7f92b94871f9f

[View certificate details](https://x509.io/?cert=MIICXzCCAgWgAwIBAgIQI5YKsEF5ScYWgpfnnjaywTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDUxOTE1NDQwMFoXDTI1MDUxODE1NDQwMFowWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkRUQyBDb21tdW5pY2F0aW9uczEsMCoGA1UEAwwjRFRDIENvbW11bmljYXRpb25zIFNIQUtFTiBDZXJ0IDA1NjIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQpOQV9Us8UmakX%2Bpsrmhfo3z%2B9p0VF1Pt4pDW7u%2F0NsFPsnwv%2BHcRanio0k2VxFIM6bbjpiempYkTWV8oEw9Wxo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwNTYyMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFLmAYVnFY5qDCg7CL34dgzbqdqegMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQDFofdODON37%2FkYZZpZ%2BEi88BmH63aTvFRDLy8sBf2qSgIgFT1PqLgYszXi8fjSSAwajaxGf5qgBL90MHRPXXQUOrs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0562', but common name is 'DTC Communications SHAKEN Cert 0562' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC