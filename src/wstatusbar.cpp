/* vi:set ts=8 sts=8 sw=8:
 *
 * Practical Music Search
 * Copyright (c) 2006-2011  Kim Tore Jensen
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

#include "window.h"
#include "console.h"
#include "curses.h"
#include <string>
#include <vector>

using namespace std;

extern vector<Logline *> logbuffer;
extern Curses curses;

void Wstatusbar::drawline(int rely)
{
	vector<Logline *>::reverse_iterator i;

	curses.wipe(rect);
	for (i = logbuffer.rbegin(); i != logbuffer.rend(); i++)
	{
		if ((*i)->level > MSG_LEVEL_INFO)
			continue;

		curses.print(rect, rely, 0, (*i)->line.c_str());
		break;
	}
}