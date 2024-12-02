import java.io.File

fun main(){

    var reports = mutableListOf<String>()
    File("day2.txt").forEachLine{
        reports.add(it)
    }

    var result = 0
    for (report in reports) {
        var reportList = report.split(" ")
        var increaseValid = true
        var decreaseValid = true

        var l = 0
        var r = 1
        var problemDampnerCount = 0
        while (r < reportList.size) {
            if (reportList[l].toInt() >= reportList[r].toInt() || reportList[l].toInt() + 3 < reportList[r].toInt()) {
                if (problemDampnerCount > 0) {
                    increaseValid = false
                    break
                }
                problemDampnerCount += 1
                l -= 1
                if (l < 0) {
                    l += 2
                    r += 1
                } else {
                    if (reportList[l].toInt() < reportList[r].toInt() && reportList[l].toInt() + 3 >= reportList[r].toInt()) {
                        l += 2
                        r += 1
                    } else {
                        increaseValid = false
                        break
                    }
                }
                
            } else {
                l += 1
                r += 1
            }
        }

        l = 0
        r = 1
        problemDampnerCount = 0
        while (r < reportList.size) {
            if (reportList[l].toInt() <= reportList[r].toInt() || reportList[l].toInt() - 3 > reportList[r].toInt()){
                if (problemDampnerCount > 0) {
                    decreaseValid = false
                    break
                }
                problemDampnerCount += 1
                l -= 1
                if (l < 0) {
                    l += 2
                    r += 1
                } else {
                    if (reportList[l].toInt() > reportList[r].toInt() && reportList[l].toInt() - 3 <= reportList[r].toInt()) {
                        l += 2
                        r += 1
                    } else {
                        decreaseValid = false
                        break
                    }
                }
            } else {
                l += 1
                r += 1
            }
        }

        

        if (increaseValid || decreaseValid ) {result += 1}
    }

    println(result)
}