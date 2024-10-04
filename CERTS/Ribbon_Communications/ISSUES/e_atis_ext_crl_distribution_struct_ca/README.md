# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications

Name: e_atis_ext_crl_distribution_struct_ca\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 18 Oct 21 00:00 UTC\
Description: STI intermediate and end-entity certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry. The DistributionPoint entry shall contain a distributionPoint field identifying the HTTP URL reference to the file containing the SHAKEN CRL hosted by the STI-PA, and a cRLIssuer field that contains the DN of the issuer of the CRL.

### Leaf Certificates

No error, warning, or notice level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | SHAKEN Ribbon Issuing CA 2 | [view](../../CERTS/2b48949bed753b3edc3645a638704bd1e708cd047117e3fe8af6cb144fa3a8c5/README.md) | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 04 Oct 24 16:29 UTC