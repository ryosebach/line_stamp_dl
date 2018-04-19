require 'open-uri'
require 'nokogiri'

def usage
prog = __FILE__
puts "Usage: #{prog} <URL>"
exit 1
end

url = ARGV.shift
startNo = ARGV.shift.to_i
usage unless url

a = Dir::entries(".")
startNo = a[a.size - 2].split("_").first.to_i
startNo = 1 if startNo == 0

lastNum = url.split("articleNo=")[1].to_i
for no in startNo..lastNum
  article_url = url.split("articleNo=").first + "articleNo=" + no.to_s
  i = 0
  doc = Nokogiri::HTML.parse(open(article_url))
  folder_name = sprintf("%03d", no).to_s + "_" +  doc.title.split('Nab').first.split(" | ")[1] + "/"
  FileUtils.mkdir_p(folder_name) unless FileTest.exist?(folder_name)
  doc.xpath('//section[@id="_comicTop"]').first.children.css('img').each do |elem|
    file_url = elem.attributes["src"].value
    filename = i.to_s + ".jpg"
	filePath = folder_name + filename
    open(file_url) do |source|
      open(filePath, "w+b") do |o|
        o.print source.read
      end
    end
    i = i + 1
  end
end
