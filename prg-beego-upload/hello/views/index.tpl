<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
 

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;

      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    header {
      padding: 100px 0;
    }
    input{
    outline-style: none ;
    border: 1px solid #ccc; 
    border-radius: 3px;
    }

    footer {
    outline-style: none ;
    border: 1px solid #ccc; 
    border-radius: 3px;
    padding: 14px 14px;
    width: 620px;
    font-size: 24px;
     box-shadow: inset 0 1px 1px rgba(0,0,0,.075),0 0 8px rgba(102,175,233,.6)
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }

  </style>
</head>

<body style="text-align:center;">
  <header>
    <h1 class="header">Welcome to POL</h1>
    <div class="description">
      You can upload file to POL server by  this button
    </div>
  </header>
  
<form id="fform" method="POST" enctype="multipart/form-data" action="/abc">   
    <input id="myfile" name="myfile" type="file" />  
	<br />
    <input type="submit" value=" 上传 "  style="width:260px" />
</form>
</body>
</html>
