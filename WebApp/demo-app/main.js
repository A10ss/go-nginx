function table1(tdcount){
    var table=document.createElement("table");
    var tr=document.createElement("tr");
    for(var i=0;i<tdcount;i++){
        var td=document.createElement("td");
        td.innerHTML=i+1;
        td.align="center"
        tr.appendChild(td);
    }
    table.appendChild(tr);
    table.width="100%";
    table.height=50;
    return table;
}
for(var i=1;i<=10;i++){
    document.getElementsByTagName("body")[0].appendChild(table1(i));
}