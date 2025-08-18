# STIR/SHAKEN CA Ecosystem Compliance

## Telonium

Name: e_atis_ext_crl_distribution_struct\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 16 Jan 22 00:00 UTC\
Description: STI intermediate and end-entity certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry. The DistributionPoint entry shall contain a distributionPoint field identifying the HTTP URL reference to the file containing the SHAKEN CRL hosted by the STI-PA, and a cRLIssuer field that contains the DN of the issuer of the CRL.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | SHAKEN 963J | [view](../../CERTS/ceef9f6d88a1efc4c056ecaf0f5d42c5f106fb5fc895c8ccc1f7ea97b6ac1093/README.md) | CRL Distribution Point shall contain a CRLIssuer field |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 18 Aug 25 21:13 UTC