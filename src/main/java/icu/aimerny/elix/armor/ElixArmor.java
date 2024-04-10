package icu.aimerny.elix.armor;

import net.minecraft.item.ArmorItem;
import net.minecraft.item.ArmorMaterial;
import net.minecraft.recipe.Ingredient;
import net.minecraft.sound.SoundEvent;

public abstract class ElixArmor implements ArmorMaterial {

    protected static final int[] BASE_DURABILITY = new int[] {13, 15, 16, 11};
    protected static final int[] PROTECTION_VALUES = new int[] {2, 4, 6, 8};

}
