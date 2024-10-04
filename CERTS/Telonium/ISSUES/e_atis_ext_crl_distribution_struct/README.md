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
| error | SHAKEN 847K | [view](../../CERTS/b973a038898de5f18bcf1e65751096cf4f788fb8f7d2cb53e29509e333600472/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | SHAKEN 963J | [view](../../CERTS/ceef9f6d88a1efc4c056ecaf0f5d42c5f106fb5fc895c8ccc1f7ea97b6ac1093/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | SHAKEN 854K | [view](../../CERTS/bfea21afc2db20c52f74b16c054d12ff6e839acc8b18401aced0154ef7e03750/README.md) | CRL Distribution Point shall contain a CRLIssuer field |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 04 Oct 24 16:29 UTC