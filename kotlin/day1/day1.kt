import kotlin.math.*
import java.io.File
import java.io.BufferedReader

fun main(args: Array<String>) {
    var listOne = mutableListOf<Int>()
    var listTwo = mutableListOf<Int>()

    File("day1.txt").forEachLine {
        val line = it.split("   ");
        listOne.add(line[0].toInt());
        listTwo.add(line[1].toInt());
    }

    val length = listOne.size
    listOne.sort()
    listTwo.sort()

    var mapOne = mutableMapOf<Int, Int>()
    var mapTwo = mutableMapOf<Int, Int>()

    var result = 0
    for (i in 0 until length) {
        result += abs(listOne[i] - listTwo[i])

        var countOne = mapOne.get(listOne[i])
        var countTwo = mapTwo.get(listTwo[i])

        if (countOne == null) {
            countOne = 0
        }
        if (countTwo == null) {
            countTwo = 0
        }

        mapOne[listOne[i]] = countOne + 1
        mapTwo[listTwo[i]] = countTwo + 1
    }

    println(result)
    var resultTwo = 0

    for ((key, value) in mapOne) {
        resultTwo += (key * (if (mapTwo[key] != null){mapTwo[key]!!}else{0})) * value
    }

    println(resultTwo)
    
}
