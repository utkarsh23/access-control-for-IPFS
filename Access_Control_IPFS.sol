pragma solidity >=0.4.22 <0.7.0;

contract AccessControlIPFS {
    struct FileData {
        address owner;
        mapping(address => bool) access;
        address[] allowedAddresses; 
    }

    mapping(bytes32 => FileData) fileMapping;

    function addBlock(address sender, bytes32 chunk, address[] memory grantees) public {
        if (chunk[0] == 0)
            return;
        FileData storage fileData;
        fileData.owner = sender;
        fileMapping[chunk] = fileData;
        address grantee;
        for (uint256 granteeInd = 0; granteeInd < grantees.length; granteeInd++) {
            grantee = grantees[granteeInd];
            fileData.access[grantee] = true;
            fileData.allowedAddresses.push(grantee);
        }
    }

    function addBlockMultiple(address sender, bytes32[] memory chunks, address[] memory grantees) public {
        bytes32 chunk;
        address grantee;
        address[] memory emptyAddresses;
        for (uint256 i = 0; i < chunks.length; i++) {
            chunk = chunks[i];
            if (chunk[0] == 0)
                continue;
            fileMapping[chunk] = FileData(sender, emptyAddresses);
            for (uint256 granteeInd = 0; granteeInd < grantees.length; granteeInd++) {
                grantee = grantees[granteeInd];
                fileMapping[chunk].access[grantee] = true;
                fileMapping[chunk].allowedAddresses.push(grantee);
            }
        }
    }

    function grantAccess(address sender, address grantee, bytes32 chunk) public {
        if (chunk[0] == 0)
            return;
        FileData storage fileData = fileMapping[chunk];
        if (sender != fileData.owner)
            return;
        fileData.access[grantee] = true;
        fileData.allowedAddresses.push(grantee);
    }

    function grantAccessMultiple(address sender, address grantee, bytes32[] memory chunks) public {
        bytes32 chunk;
        for (uint256 i = 0; i < chunks.length; i++) {
            chunk = chunks[i];
            if (chunk[0] == 0)
                continue;
            FileData storage fileData = fileMapping[chunk];
            if ((i == 0) && (sender != fileData.owner))
                return;
            fileData.access[grantee] = true;
            fileData.allowedAddresses.push(grantee);
        }
    }

    function removeAccess(address sender, address grantee, bytes32 chunk) public {
        if (chunk[0] == 0)
            return;
        FileData storage fileData = fileMapping[chunk];
        if (sender != fileData.owner)
            return;
        fileData.access[grantee] = false;
    }

    function removeAccessMultiple(address sender, address grantee, bytes32[] memory chunks) public  {
        bytes32 chunk;
        for (uint256 i = 0; i < chunks.length; i++) {
            chunk = chunks[i];
            if (chunk[0] == 0)
                continue;
            FileData storage fileData = fileMapping[chunk];
            if ((i == 0) && (sender != fileData.owner))
                return;
            fileData.access[grantee] = false;
        }
    }

    function checkAccess(address grantee, bytes32 chunk) public view returns (bool hasAccess) {
        hasAccess = false;
        if (chunk[0] == 0)
            return hasAccess;
        FileData storage fileData = fileMapping[chunk];
        if ((fileData.owner == grantee) || (fileData.access[grantee] == true)) {
            hasAccess = true;
            return hasAccess;
        }
    }

    function checkAccessMultiple(address grantee, bytes32[] memory chunks) public view returns (bool[] memory _checkAccessMultiple) {
        _checkAccessMultiple = new bool[](chunks.length);
        bytes32 chunk;
        for (uint256 i = 0; i < chunks.length; i++) {
            chunk = chunks[i];
            _checkAccessMultiple[i] = false;
            if (chunk[0] == 0)
                continue;
            FileData storage fileData = fileMapping[chunk];
            if ((fileData.owner == grantee) || (fileData.access[grantee] == true)) {
                _checkAccessMultiple[i] = true;
            }
        }
    }

    function deleteBlock(address sender, bytes32 chunk) public {
        if (chunk[0] == 0)
            return;
        delete fileMapping[chunk];
    }

    function deleteBlockMultiple(address sender, bytes32[] memory chunks) public {
        bytes32 chunk;
        for (uint256 i = 0; i < chunks.length; i++) {
            chunk = chunks[i];
            if (chunk[0] == 0)
                continue;
            delete fileMapping[chunk];
        }
    }
}
