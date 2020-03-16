#最简单的http文件上传  文件名 up.rb ,  执行 ruby up.rb 
require "webrick"
$port = 8886
class PostSampleServlet < WEBrick::HTTPServlet::AbstractServlet
  def initialize(server, limit)
    @max_content_length = limit
    super
  end

  def do_GET(req, res)
    content_length = req['content-length'].to_i
    if content_length > @max_content_length
      raise WEBrick::HTTPStatus::BadRequest, "body is too large"
    end
    if data = req.query["data"]
      filename = data.filename
      puts filename
    end
    res.body =<<-_end_of_html_
<html>
 <form method="POST" enctype="multipart/form-data">
  <input type="file" name="data" /><input type="submit" /></form>
 filename = #{WEBrick::HTMLUtils.escape(filename.inspect)}
 <pre>#{ 
  if filename 
    p filename
    a=File.new(filename,'wb')
    a.write data
    a.close
    "#{data.size} 上传完成    #{Dir.pwd }/#{filename}"
  end
 }</pre>
</html>
    _end_of_html_
    res["content-type"] = "text/html"
  end

  def do_POST(req, res)
    do_GET(req, res)
  end
end

svr = WEBrick::HTTPServer.new(:Port=> $port)
svr.mount("/", PostSampleServlet, 50000000) #50MB
trap(:INT){ svr.shutdown }
svr.start
