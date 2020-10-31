pragma solidity  ^0.6.6;


contract SeedManagement {


    struct Seed{
        address spa;
        string LotNumber;
        string owner;
        string crop;
        string variety;
        string SourceTagNo;
        string SourceClass;
        string DestinationClass;
        string SourceQuantity;
        string SourceDateOfIssue;
        string growerName;
        string spaName;
        string dateOfIssue;
        string sourceStorehouse;
        string destiStorehouse;
        string sgName;
        string sgId;
        string FinYear;
        string Season;
        string landRecordsKhatano;
        string landRecordsPlotno;
        string landRecordsArea;
        string cropRegCode;
        string SppName;
        string SppID;
        string TotalQuantityProcessed;
        string processingDate;
        string verificationDate;
        string SampleSecretCode;
        string SamplePassed;
        string SampleTestDate;
        string CertificateNo;
        string CerticateDate;
        string TagSeris;
        string TagIssuedRangeFrom;
        string TagIssuedRangeTo;
        string NoOfTagsIssued;
        string CertificateValidityInMonth;
    }

    mapping(string => Seed) public seeds;

    event seedCreation(string indexed _lotNumber, string indexed Season , string indexed dateOfIssue );
    event seedCertification(string indexed CertificateNo, string indexed CerticateDate, string indexed CertificateValidityInMonth);
    event seedTesting(string indexed SampleSecretCode, string indexed SamplePassed, string indexed SampleTestDate);



    //Addition of new product to blockchain
    function createSeed(string memory _lotNumber, string  memory _owner,  string memory _crop,  string memory _variety, string memory _sourceTagNo ,string memory _sourceClass, string memory _destinationClass, string  memory _sourceQuantity,  string memory _sourceDateOfIssue ,string memory _growerName ,string memory _spaName, string memory _dateOfIssue ,string memory _sourceStorehouse, string memory _destiStorehouse ,string memory _sgName, string memory _sgId ,string  memory _finYear, string memory _season ,string memory _landRecordsKhatano , string  memory _landRecordsPlotno , string memory _landRecordsArea , string memory _cropRegCode , string memory _sppName ,string  memory _sppID) public
    {

      require(seeds[_lotNumber] == "", "Product was already registered");

      Seed memory new_seed = Seed(msg.sender,  _lotNumber,  _owner,  _crop,   _variety,   _sourceTagNo ,  _sourceClass,   _destinationClass,    _sourceQuantity,    _sourceDateOfIssue ,  _growerName ,  _spaName,   _dateOfIssue ,  _sourceStorehouse,   _destiStorehouse ,  _sgName,  _sgId , _finYear,  _season , _landRecordsKhatano , _landRecordsPlotno ,  _landRecordsArea , _cropRegCode , _sppName , _sppID ,"" ,"" ,"" ,"" ,"" ,"" ,"" ,"","","","","","",true);
      seeds[_lotNumber ] = new_seed;

      emit seedCreation(_lotNumber, _season,  _dateOfIssue);
    }

    //Called when product is imported, exported or purchased
    function certify(string memory _productId , string memory _from , string memory _to, string memory _type) public
    {
        require(products[_productId].exists == true, "Product does not exist");

        emit productTransfer(_from , _to , _productId,_type);

    }

    function testify(string memory _productId , string memory _from , string memory _to, string memory _type) public
    {
        require(products[_productId].exists == true, "Product does not exist");

        emit productTransfer(_from , _to , _productId,_type);

    }


    //Called to check the existence of product
    function checkExistence(string memory _productId) public view returns(bool)
    {
    //   require(products[_productId].exists != true, "Product does not exist");
    if(products[_productId].exists == false)
    {
        return false;
    }
    else
      {
          return true;
      }
    }
}
