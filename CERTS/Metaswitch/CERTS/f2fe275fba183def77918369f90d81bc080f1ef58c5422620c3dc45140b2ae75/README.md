# STIR/SHAKEN CA Ecosystem Compliance

## Certificate American Broadband SHAKEN Cert 355D

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 83 day(s)\
Subject: CN=American Broadband SHAKEN Cert 355D, O=American Broadband, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/f786e53509092c8a45b19b7dcd6886b5316d333b

[View certificate details](https://understandingwebpki.com/?cert=MIICXzCCAgWgAwIBAgIQXb5l%2FXT4Re1URPM7XE5hSDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUwNTIxMDI1NFoXDTI0MDUwNDIxMDI1NFowWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkFtZXJpY2FuIEJyb2FkYmFuZDEsMCoGA1UEAwwjQW1lcmljYW4gQnJvYWRiYW5kIFNIQUtFTiBDZXJ0IDM1NUQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASHDsAID3%2Bu%2F7Xv5AU9r6PJ8FQi5rayLQAgj7zJko0jKnzdkrDpq1fjE%2BluclcipIXP3Pq%2F16jR7OoNmGDSwFWVo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQzNTVEMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFL5xqjVymgE1zuEO5Jjv1%2FmSM50lMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIFtCzQhaLeaRYSGMvb59ioNELqw73bkpdk7DMVy5ttZ6AiEAkgfVA%2B4Htu8VNw2MAPmMkor9hXIIhlfrsHInobGlsQA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC