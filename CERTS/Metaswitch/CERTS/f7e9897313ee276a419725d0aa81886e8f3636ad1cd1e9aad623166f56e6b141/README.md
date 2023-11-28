# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Yelcot SHAKEN Cert 1733

Tested At: 28 Nov 23 15:59 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 110 day(s)\
Subject: CN=Yelcot SHAKEN Cert 1733, O=Yelcot, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/02f31fd9be31f3b5cc9fa396dc9fab04f1dc93e1

[View certificate details](https://understandingwebpki.com/?cert=MIICRjCCAe2gAwIBAgIQG5At462VYEBclL2GUgCReTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxNzE3MjU0N1oXDTI0MDMxNjE3MjU0N1owQDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlllbGNvdDEgMB4GA1UEAwwXWWVsY290IFNIQUtFTiBDZXJ0IDE3MzMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAShxrcyN%2FMVpRh0cbuaAAAIuYKBw3SyStFsAbehoEjPTY2Aezb%2B%2FEHVQLrLwnzFn2BbR4O42w053ruwp539VQxvo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQxNzMzMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFC0Oj9LeGCoONM5mPKLBf8pmm79zMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIBcWk7Iyi3zduVFzoqYfwNfuS%2FpIqP1dwVhnVKzfTXmQAiByjvSl7Z8%2BcBl6dq6snGSt%2B1ZbHg97UOWZUe7JIT%2BNcw%3D%3D)

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


Generated: 28 Nov 23 16:15 UTC